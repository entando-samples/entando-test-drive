---
slug: 03-entando
id: zmjfcpmjxpfs
type: challenge
title: Creating a Page Template
teaser: Prepare a Page Template with Two Frames
notes:
- type: text
  contents: |-
    **Page Templates** provide the scaffolding of a Page and are constructed using two main elements:
    1) A JSON Configuration
    2) A Template

    See the next few notes for details...
- type: text
  contents: |-
    # JSON Configuration
    This field lists the set of frames that can be used on a page. Each item in the frames array represents a frame or slot in the page, characterized by the following values:
     - ```pos``` - a zero-based position index (starts from zero). This value is used in APIs to address a specific widget on the page.
     - ```descr``` - the frame description displayed on Page Design view
     - ```mainFrame``` - designates the primary frame in the Page Template
     - ```defaultWidget``` - widget code for a default widget to use in this frame. Page Template developers can use this field to provide suggestions on common widgets, e.g. header and footer widgets.
     - ```sketch``` - an object with 4 coordinates (x1,x2,y1,y2) to allow the developer to place the widgets in the Page preview. Sketch’s x and y values go from 0 to 11 (similar to columns in Bootstrap), so if you want to place a 2x2 frame at the top left corner of the page, the values would be ```x1: 0, x2: 1 y1: 0 y2: 1```.
- type: text
  contents: |-
    # Template
    This uses Freemarker code to setup the HTML page itself.
    - To add a frame in a specific place of the page, add ```<@wp.show frame=0 />```, where frame is the pos variable from the ***JSON configuration***. ```<#assign wp=JspTaglibs["/aps-core"]>``` is required at the top of the template to setup the wp variable.
    - Common code can be shared across pages by using UX Fragments and ```<@wp.fragment code="\<FRAGMENT\_CODE\>" escapeXml=false /\>```.
tabs:
- title: App Builder
  type: website
  path: /app-builder/
  url: https://kubernetes-vm.${_SANDBOX_ID}.instruqt.io/app-builder/
- title: shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 1200
---
1. In the Left Sidebar, Click on *Pages > Templates*
- Click ***ADD*** in the bottom right corner
---
2. Enter the following details:

- **Code:** double_frame

- **Name:** Double Frame

- **JSON Configuration:**

```
{
  "frames": [
    {
      "pos": 0,
      "descr": "Frame 1",
      "mainFrame": false,
      "defaultWidget": null,
      "sketch": {"x1": 0, "y1": 0, "x2": 11, "y2": 1}
    },
    {
      "pos": 1,
      "descr": "Frame 2",
      "mainFrame": false,
      "defaultWidget": null,
      "sketch": {"x1": 0, "y1": 2, "x2": 11, "y2": 3}
    }
  ]
}
```
- **Template:**
```
<#assign wp=JspTaglibs["/aps-core"]>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<html>
  <head>
      <title><@wp.currentPage param="title" /></title>
  </head>
  <body>
    <h1><@wp.currentPage param="title" /></h1>
    <div><@wp.show frame=0 /></div>
    <div><@wp.show frame=1 /></div>
  </body>
</html>
```

Include the following fragment in the head section if you want to make use of the user's identity information from Keycloak:
```
<@wp.fragment code="keycloak_auth" escapeXml=false />
```
---
3. You should be able to see the ***Template Preview*** showing the desired two-frame layout
---
4. Click ***Save*** to save your template.
---