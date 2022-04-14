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
    **INSTALLING PBCs INTO THE APP BUILDER**

    In this section, we will learn how to add a newly developed module to a web application.  Modules can be referred to as, **Packaged Business Capabilities** or PBCs.

    We will install a Tableau PBC into our application on a new page.
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

**STEP 1: ADD THE NEW  PBC INTO THE LOCAL REPOSITORY**

A **Creator** (an individual developer or a developer team) has created a new Tableau Packaged Business Capability (**PBC**) displaying SAP Returns. Your goal as a **Composer** is to add this to the current application. <u>(Ensure you are on the Shell tab on the left.)</u>

**<u>Run</u>** the following command to deploy the Tableau PBC into our local **Component Repository** (which holds all modules which Composers can use to deploy into apps).


```
ent bundler from-git -r https://github.com/entando-samples/tableau-example-bundle -d | ent kubectl apply -n entando -f -
```

This command will pull the module's code from **Github** and configure new resources in **Kubernetes** allowing the module to now be used within our **App Builder**.

**STEP 2:  INSTALL THE PBC INTO OUR APP BUILDER**

On the left, **<u>click on the App Builder tab</u>**,

then **<u>click on the "Don't show next time</u>** checkbox, and

then **<u>click Close </u>** and **<u>click Yes</u>** to end the wizard.

In the Left Sidebar, **<u>click Repository.</u>**

As shown below, inside the tableaujs-bundle card, **<u>click on the Install button.</u>**

![Image Description](../assets/tableau_bundle.png)

As shown below, after the installation process is finished, the install button will switch to an **uninstall** green button.

![Image Description](../assets/installedSuccessfully.png)

Note, this process can be used to update or downgrade the version of any specific PBC.

**STEP 3: VERIFY THE PBC IS AVAILABLE**

In the Left Sidebar, **<u>Click on Components > Micro Frontends & Widgets.</u>**

All out-of-the-box and user-installed components and PBCs are listed here.   Scroll down to the **User** section.   Observe that your tableaujs-bundle is displayed here and is now available to add to various pages in your application.

![Image Description](../assets/list_widgets.png)

Congratulations! You have completed this challenge.  In the next challenge, you will learn how to create a page which includes this PBC.

When you are ready, validate this challenge by clicking **<u>Check</u>**.
