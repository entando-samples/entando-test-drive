---
slug: 01-entando
id: wwtqkv7pvfl9
type: challenge
title: 'Welcome Wizard: Creating a Page'
teaser: 'Welcome Wizard: Creating a Page containing content and other Entando components'
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
- title: shell
  type: terminal
  hostname: named-k3s
- title: Entando Website
  type: website
  path: /app-builder/
  url: https://dev.entando.org/
- title: Authentication
  type: external
  url: https://named-k3s.${_SANDBOX_ID}.instruqt.io/app-builder/
- title: App Builder
  type: website
  url: https://named-k3s.${_SANDBOX_ID}.instruqt.io/app-builder/
difficulty: basic
timelimit: 120000
---
Click on `Authentication` tab to authenticate. It will open the login screen in new browser tab.

Once you see the login screen, enter the following login details:
- Username: admin
- Password: adminadmin


Change the password if required.

Click on `App Builder` tab inside instruqt and reload. It will show App Builder dashboard.

`Follow the Welcome Wizard. On completion of wizard it will open the newly created page in the external browser as shown below`

![Image Description](../assets/hello_world_app.png)

To complete this challenge, press **Check**.
