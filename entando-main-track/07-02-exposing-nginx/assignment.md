---
slug: 02-exposing-nginx
id: olmaf9azolxk
type: challenge
title: Expose the NGINX service
teaser: Use a NodePort to expose the NGINX webserver
notes:
- type: text
  contents: You've just deployed NGINX. To view the deployment, create a NodePort
    service for it.
difficulty: basic
timelimit: 600
---
Use the terminal to add a NodePort service:

```
kubectl apply -f service.yml
```