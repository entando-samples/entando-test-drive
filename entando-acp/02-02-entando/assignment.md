---
slug: 02-entando
id: phlghnumsrab
type: challenge
title: 'S&P Returns Dashboard: Install the PBC'
teaser: In the previous challenge, we learned how an Application Composition Platform
  works. In this challenge we explore some advanced ACP capabilities to create a dashboard
  Page.
notes:
- type: text
  contents: |-
    # Compose an S&P Financial Returns Dashboard page
    Our users (Consumers) have asked us to give them one page where they can easily see S&P Financial Returns and Economic Indicators.
    Fortunately, the Entando ACP provides a way to use existing components wrapped into Packaged Business Capabilities (PBC) and compose them to create this new page.
    This PBC was created by our Creator team and validated by our Curator. That means we can use it in our application.
    You are the main Composer of this challenge. Please proceed to the PBC installation to finish it.

    A video with further explanation is waiting for you in the next slide.
- type: video
  url: https://www.youtube.com/embed/W5Xnv5Uwrxw
tabs:
- title: shell
  type: terminal
  hostname: named-k3s
- title: App Builder
  type: website
  url: https://named-k3s.${_SANDBOX_ID}.instruqt.io/app-builder/
difficulty: basic
timelimit: 1200
---
Hi Consumer,
Our Curator has validated the Tableau PBC and it is now available.

Please run the following command to deploy it in our Component Repository, where all components used in the application are listed.

```
ent bundler from-git -r https://github.com/entando-samples/tableau-example-bundle -d | ent kubectl apply -n entando -f -
```

Once it's registered, we need to install it from our Entando Component Repository (ECR). The installation process configures new resources on Kubernetes to run our components.

To execute that task, be sure to be on the **App Builder tab**.

In the Left Sidebar, Click on **Repository**.

Click on the install button.

![Image Description](../assets/tableau_bundle.png)

When the installation process is finished, the install button switches to an uninstall plain green button.

![Image Description](../assets/installedSuccessfully.png)

---

You can navigate to the components entry list to see newly installed components from the PBC.

In the Left Sidebar, Click on **Components** > **Micro Frontends & Widgets**.

All the micro frontends and widgets are listed here. If some are available out of the box, the installed components from the PBC are available under the **User** section.

The **Tableau Sample Widget** is here.

![Image Description](../assets/list_widgets.png)

---

Congratulations! You completed this challenge in which you've installed your first PBC.

The next challenge will learn you how to create a page and compose it with the component included in the PBC.
When you are ready, validate this challenge by clicking **Check**.

