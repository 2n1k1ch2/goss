package suspicion

import "goss/pkg/capture"

func ruleBigCluster(_ *capture.Cluster, obj *capture.Object) {
	switch obj.Count {
	case 100:
		obj.Score += 5
	case 500:
		obj.Score += 10
	case 1000:
		obj.Score += 20
	}
}

func ruleSendNoRecv(cluster *capture.Cluster, obj *capture.Object) {
	if obj.Status == capture.CHAN_SEND {
		i := 0
		for _, v := range *cluster {
			if v.Status == capture.CHAN_RECEIVE {
				i++
			}
		}
		if i == 0 {
			obj.Score += 10
		}
	}
}
func ruleSleep(_ *capture.Cluster, obj *capture.Object) {

}
