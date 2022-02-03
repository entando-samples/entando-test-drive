---
slug: 02-entando
id: dcbvpxyt6jxx
type: challenge
title: Tableau bundle installation
teaser: Add MFE Bundle and see it as widget on page
notes:
- type: text
  contents: |
    # Tableau bundle installation
    This Challenge will explain you the bundle **installation**.
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
# Install the tableau js bundle

Execute below command:
```
ent bundler from-git -r https://github.com/nshaw/2021-12-tableau-bundle.git -d | ent kubectl apply -n entando -f -
```

Goto App Builder tab > goto Repository tab on side bar of application. You can see MFE is ready to install.
![Image Description](../assets/tableau_bundle.png)

# Click on Install button & wait for successful installation.

