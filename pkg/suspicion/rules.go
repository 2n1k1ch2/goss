package suspicion

import "goss/pkg/cluster"

func ruleBigCluster(_ *cluster.Cluster, obj *cluster.Object) {
	switch obj.Count {
	case 100:
		obj.Score += 5
	case 500:
		obj.Score += 10
	case 1000:
		obj.Score += 20
	}
}

func ruleSendNoRecv(cl *cluster.Cluster, obj *cluster.Object) {
	if obj.Status == cluster.CHAN_SEND {
		i := 0
		for _, v := range *cl {
			if v.Status == cluster.CHAN_RECEIVE {
				i++
			}
		}
		if i == 0 {
			obj.Score += 10
		}
	}
}
func ruleSleep(_ *cluster.Cluster, obj *cluster.Object) {

}
