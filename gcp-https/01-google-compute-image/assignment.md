---
slug: google-compute-image
id: sbykcwvmziaw
type: challenge
title: Google Compute Image
notes:
- type: text
  contents: |-
    <iframe width="900px" height="100px" src="https://tools.arantius.com/stopwatch" >
    </iframe><iframe width="900px" height="350px"
    src="https://www.youtube.com/embed/EvbFHywvr0g">
    </iframe>
  url: https://www.youtube.com/embed/EvbFHywvr0g
- type: text
  contents: |-
    <iframe width="900px" height="350px" src="https://tools.arantius.com/stopwatch" >
    </iframe>
tabs:
- title: Shell
  type: terminal
  hostname: entando-vm
- title: Editor
  type: code
  hostname: entando-vm
  path: /root/
- title: Keycloak
  type: website
  url: http://entando-vm.${_SANDBOX_ID}.nip.io/app-builder/
  new_window: true
- title: App Builder
  type: website
  url: https://entando-vm.${_SANDBOX_ID}.nip.io/app-builder/
- title: Static GCP AppBuilder
  type: website
  url: https://34.125.13.6.nip.io/app-builder/
difficulty: basic
timelimit: 1200000
---
Use the terminal to execute shell script:

```
sudo ./start.sh
```

``
For Every pass phrase use same password : entando
``

Wait for 5-10 mins then check for pods:
```
sudo kubectl -n entando get pods
```

Watch Entando startup. It can take around 10 minutes before the application is fully deployed and ready.
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
Click on `keycloak` tab to authenticate. It will open the login screen in new browser tab.

Once you see the login screen, enter the following login details:
- Username: admin
- Password: adminadmin


Change the password if required.

Click on `App Builder` tab inside instruqt and reload. It will show App Builder dashboard.

``Follow the Welcome Wizard. On completion of wizard it will open the newly created page in the external browser as shown below``

![Image Description](../assets/hello_world_page.png)

To complete this challenge, press **Check**.