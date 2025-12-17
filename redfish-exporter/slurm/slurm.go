package slurm

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var (
	defaultSlurmUsername   = "root"
	slurmRestClientTimeout = time.Minute * 5
	maxRetries             = 5
	baseDelay              = 1 * time.Second
)

// Retryable is a function type that represents any API call function you we want to retry
type Retryable func() (interface{}, *http.Response, error)

// CallWithRetry is a generic retry function that takes a function, retry count, and delay
func CallWithRetry(call Retryable, maxRetries int, baseDelay time.Duration) (interface{}, *http.Response, error) {
	var resp *http.Response
	var err error
	var res interface{}

	for i := 0; i < maxRetries; i++ {
		res, resp, err = call()
		if err == nil && resp.StatusCode == http.StatusOK {
			return res, resp, nil // success
		}

		// exponential backoff delay
		delay := time.Duration(math.Pow(2, float64(i))) * baseDelay
		log.Printf("attempt %d: Error: %v, retrying in %v\n", i+1, err, delay)
		time.Sleep(delay)
	}

	return nil, nil, fmt.Errorf("after %d retries, error: %w", maxRetries, err)
}

type SlurmServerConfig struct {
	URL         string
	Username    string
	BearerToken string
}

func DrainNodeWithScontrol(nodeName, reason, excludeStr, scontrolPath string) error {

	if excludeStr != "" {
		curReason, err := GetNodeReasonWithScontrol(nodeName, scontrolPath)
		if err != nil {
			log.Printf("GetNodeReasonWithScontrol returned err: %v\n", err)
			return err
		}

		if curReason != "" {
			re := regexp.MustCompile(excludeStr)
			match := re.FindAllString(curReason, -1)

			if len(match) != 0 {
				log.Printf("exlcudStr: %v, curReason: %v", excludeStr, curReason)
				log.Printf("match: %v | len: %v", match, len(match))
				return fmt.Errorf("%s: not draining node: %s | current reason: %s", ExlcudeReasonSet, nodeName, curReason)
			}
		}
	}
	cmd := fmt.Sprintf("%s update NodeName=%s State=DRAIN Reason=\"%s\"", scontrolPath, nodeName, reason)
	res := LocalCommandOutput(cmd)
	log.Printf("Drain node result: %s", res)
	return nil
}

func GetNodeReasonWithScontrol(nodeName, scontrolPath string) (string, error) {
	type scontrolShowNode struct {
		Nodes []struct {
			Reason string `json:"reason"`
		} `json:"nodes"`
	}

	cmd := fmt.Sprintf("%s show node %s --json", scontrolPath, nodeName)
	ret := LocalCommandOutput(cmd)

	if ret == "" {
		return "", fmt.Errorf("failed to get current node reason")
	}

	res := scontrolShowNode{}
	err := json.Unmarshal([]byte(ret), &res)
	if err != nil {
		return "", err
	}
	if len(res.Nodes) != 1 {
		return "", fmt.Errorf("show node failed for %s", nodeName)
	}
	log.Printf("get node reasons(%s): %+v\n", nodeName, res.Nodes[0].Reason)
	return res.Nodes[0].Reason, nil
}

// LocalCommandOutput runs a command on a node and returns output in string format
func LocalCommandOutput(command string) string {
	log.Printf("Running cmd: %s\n", command)
	out, _ := exec.Command("bash", "-c", command).CombinedOutput()
	return strings.TrimSpace(string(out))
}
