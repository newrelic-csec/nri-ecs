package testdata

// IntegrationOutput is the resulting payload generated by
// `integration.Publish()`
var IntegrationOutput string = `
{
	"name":                "com.newrelic.ecs",
	"protocol_version":    "3",
	"integration_version": "1.3.1",
	"data": [
		{
			"entity": {
				"name": "cluster/ecs-local-cluster",
				"type": "arn:aws:ecs:us-west-2:111111111111",
				"id_attributes":[]
			},
			"metrics": [
				{
					"clusterName": "ecs-local-cluster",
					"event_type":  "EcsClusterSample",
					"arn":         "arn:aws:ecs:us-west-2:111111111111:cluster/ecs-local-cluster"
				}
			],
			"inventory": {
				"cluster": {
					"name": "ecs-local-cluster",
					"arn":  "arn:aws:ecs:us-west-2:111111111111:cluster/ecs-local-cluster"
				}
			},
			"events": []
		},
		{
			"metrics": [],
			"inventory": {
				"host": {
					"ecsClusterName": "ecs-local-cluster",
					"ecsClusterArn":  "arn:aws:ecs:us-west-2:111111111111:cluster/ecs-local-cluster",
					"awsRegion": "us-west-2",
					"ecsLaunchType": "ec2"
				}
			},
			"events": []
		}
	]
}
`
