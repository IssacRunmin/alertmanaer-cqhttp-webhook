## Alertmanager Dingtalk Webhook

Webhook service support send Prometheus 2.0 alert message to Dingtalk.

## How To Use
Assume cqhttp working on localhost at port 5600, start this webhook on 5602
```
cd cmd/webhook
go build
go run webhook.go -defaultRobot="http://127.0.0.1:5600/send_private_msg?access_token=xxxx&user_id=xxx" -port=5602
```

* -defaultRobot: default dingtalk webhook url, all notifaction from alertmanager will direct to this webhook address.

Or you can overwrite by add annotations to Prometheus alertrule to special the dingtalk webhook for each alert rule.

```
groups:
- name: hostStatsAlert
  rules:
  - alert: hostCpuUsageAlert
    expr: sum(avg without (cpu)(irate(node_cpu{mode!='idle'}[5m]))) by (instance) > 0.85
    for: 1m
    labels:
      severity: page
    annotations:
      summary: "Instance {{ $labels.instance }} CPU usgae high"
      description: "{{ $labels.instance }} CPU usage above 85% (current value: {{ $value }})"
      cqRobot="http://127.0.0.1:5602/send_group_msg?access_token=xxxx&group_id=xxxaccess_token=xxxx&group_id=xxx"
```
