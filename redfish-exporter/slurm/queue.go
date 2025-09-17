package slurm

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/nod-ai/ADA/redfish-exporter/metrics"
)

const (
	Drain            = "DrainNode"
	ExlcudeReasonSet = "DRAIN_EXCLUDE_REASON_SET"
)

type AddEventReq struct {
	RedfishServerIP string
	SlurmNodeName   string
	Severity        string
	Action          string
	DrainReason     string
	MessageId       string
	Message         string
	ExcludeStr      string
	ScontrolPath    string
}

type eventsActionReq struct {
	redfishServerIP string
	slurmNodeName   string
	severity        string
	action          string
	drainReason     string
	messageId       string
	message         string
	excludeStr      string
	scontrolPath    string
}

type SlurmQueue struct {
	ctx   context.Context
	queue chan *eventsActionReq
}

func InitSlurmQueue(ctx context.Context) *SlurmQueue {
	return &SlurmQueue{ctx: ctx, queue: make(chan *eventsActionReq)}
}

func (q *SlurmQueue) Add(evt AddEventReq) {
	q.queue <- &eventsActionReq{
		redfishServerIP: evt.RedfishServerIP,
		slurmNodeName:   evt.SlurmNodeName,
		severity:        evt.Severity,
		drainReason:     evt.DrainReason,
		excludeStr:      evt.ExcludeStr,
		scontrolPath:    evt.ScontrolPath,
	}
}

func (q *SlurmQueue) ProcessEventActionQueue() {
	log.Println("Starting slurm queue")
	for {
		select {
		case <-q.ctx.Done():
			log.Printf("Context done, stopping slurm queue processing")
			close(q.queue)
			return
		case actionReq := <-q.queue:
			log.Printf("Processing events action req from slurm queue: %v", actionReq)
			if err := q.performEventAction(actionReq); err != nil {
				metrics.SlurmAPIFailureMetric.WithLabelValues(
					actionReq.redfishServerIP,
					actionReq.slurmNodeName,
					actionReq.severity,
					"Drain").Inc()
			} else {
				metrics.SlurmAPISuccessMetric.
					WithLabelValues(
						actionReq.redfishServerIP,
						actionReq.slurmNodeName,
						actionReq.severity,
						"Drain").Inc()
			}
		}
	}
}

func getDrainReasonString(prefix, msg, msgId, severity string) string {
	ret := fmt.Sprintf("%s:redfishlistener:%s:%s:%s", prefix, severity, msgId, msg)
	return ret
}

func (q *SlurmQueue) performEventAction(req *eventsActionReq) error {
	if len(strings.TrimSpace(req.slurmNodeName)) == 0 {
		return nil
	}

	err := DrainNodeWithScontrol(req.slurmNodeName, req.drainReason, req.excludeStr, req.scontrolPath)
	if err != nil {
		if strings.Contains(err.Error(), ExlcudeReasonSet) {
			log.Printf("Node not drained: %v", err.Error())
			return nil
		}
		log.Printf("Error draining node: %v", err)
		return err
	}

	log.Printf("Performed action: %v on slurm node: %v successfully", req.action, req.slurmNodeName)
	return nil
}
