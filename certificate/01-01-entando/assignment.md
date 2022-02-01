---
slug: 01-entando
id: k7d4boddhzlb
type: challenge
title: Setting Up the Entando Platform
teaser: Bringing up and running entando platform on Kubernetes
notes:
- type: text
  contents: |
    **Entando is the leading modular application platform for Kubernetes.**

    In this track, we will explore the application and do some basic tasks.

    ## Please go through next slides to know more about Entando platform
- type: text
  contents: |-
    <div class="box-content">
        <h1>
            <p>A Modern Platform for Modern UX</p>
        </h1>
    </div>

    <div>
        <img src="https://www.entando.com/ROOT/cmsresources/cms/images/Entando_Maerketecture_d0.svg" alt="Entando_Maerketecture.svg">
    </div>
  url: https://www.entando.com/ROOT/cmsresources/cms/images/Entando_Maerketecture_d0.svg
- type: video
  url: https://www.youtube.com/embed/EvbFHywvr0g
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: Entando Website
  type: website
  path: /app-builder/
  url: https://dev.entando.org/
- title: keycloak
  type: website
  url: https://kubernetes-vm.${_SANDBOX_ID}.instruqt.io/app-builder/
  new_window: true
- title: App Builder
  type: website
  url: https://kubernetes-vm.${_SANDBOX_ID}.instruqt.io/app-builder/
- title: Editor
  type: code
  hostname: kubernetes-vm
  path: /root/
- title: Reports
  type: service
  hostname: kubernetes-vm
  path: /
  port: 7800
  new_window: true
difficulty: basic
timelimit: 120000
---
<!-- Use the terminal to execute shell script:

```
sudo ./start.sh
``` -->
<!--
``
For Every pass phrase use same password : entando
`` -->

Wait for 5-10 mins then check for pods:
```
sudo kubectl -n entando get pods
```

Watch Entando startup. It can take around 10-15 minutes before the application is fully deployed and ready.
```
sudo kubectl -n entando get pods --watch
```

Use the terminal to check for the Entando ingresses using
```
sudo kubectl describe ingress -n entando
```

Use the terminal to get the Entando ingresses using
```
sudo kubectl get ingress -n entando
```

Use the terminal to manually patch the eci-ingress next, once the ingresses are in place
```
sudo ./patch-cm.sh
```

It can take around 3 minutes before the pod `quickstart-cm-deployment` restarted and ready.
![Image Description](../assets/cm-deployment-pod.png)

Click on `keycloak` tab to authenticate. It will open the login screen in new browser tab.

Once you see the login screen, enter the following login details:
- Username: admin
- Password: adminadmin


Change the password if required.

Click on `App Builder` tab inside instruqt and reload. It will show App Builder dashboard.

``Follow the Welcome Wizard. On completion of wizard it will open the newly created page in the external browser as shown below``

![Image Description](../assets/hello_world_page.png)

To complete this challenge, press **Check**.