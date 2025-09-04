package suspicion

import "goss/pkg/capture"

var sliceRules = []func(cluster *capture.Cluster, object *capture.Object){}

func RegisterRules(rule func(cluster *capture.Cluster, obj *capture.Object)) {
	sliceRules = append(sliceRules, rule)
}
func ScoreObject(cluster *capture.Cluster, obj *capture.Object) {
	for _, r := range sliceRules {
		r(cluster, obj)
	}
}

func init() {
	RegisterRules(ruleBigCluster)
	RegisterRules(ruleSendNoRecv)
	RegisterRules(ruleSleep)
}
