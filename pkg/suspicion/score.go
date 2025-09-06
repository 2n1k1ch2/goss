package suspicion

import "goss/pkg/cluster"

var sliceRules = []func(cluster *cluster.Cluster, object *cluster.Object){}

func RegisterRules(rule func(cluster *cluster.Cluster, obj *cluster.Object)) {
	sliceRules = append(sliceRules, rule)
}
func ScoreObject(cluster *cluster.Cluster, obj *cluster.Object) {
	for _, r := range sliceRules {
		r(cluster, obj)
	}
}

func init() {
	RegisterRules(ruleBigCluster)
	RegisterRules(ruleSendNoRecv)
	RegisterRules(ruleSleep)
}
