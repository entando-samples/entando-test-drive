---
slug: 04-entando
id: v1wubyfv891x
type: challenge
title: Creating a Widget
teaser: Learn the basics of creating an Entando widget and placing it on a page
notes:
- type: text
  contents: |-
    In this example, you will use the Entando App Builder to build and display a simple **widget** on a page.

    In a production system or a larger development environment you would build and deploy widgets differently, however this example provides a quick idea of the building blocks.
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: App Builder
  type: service
  hostname: kubernetes-vm
  path: /app-builder/
  port: 80
  new_window: true
difficulty: basic
timelimit: 1200
---
## 1. In the App Builder menu, go to: ***Components > Micro frontends & Widgets***
- At the bottom of the page, select ***ADD***
---
## 2. Create a Widget with the sample HTML code
- Enter the following details in the right sidebar:
  - **Code:** ```MyHelloWorld```
  - **Title:** ```Hello World```
  - **Title:** ```Ciao Mondo```
- In the **Custom UI field:**, enter ```<h2>Hello World</h2>```

- Click the ***SAVE*** button in the top right corner

> Note: the Custom UI Field is a freemarker template where you can put raw html and include freemarker logic. This allows you to import javascript, css, or any normal HTML.

---

## 3. Select a new Home Page
- Go to ***Pages > Settings***
- In the ***Home Page*** dropdown menu, select ***Home / Service***
- Click ***SAVE***

---

## 4. Place the widget on the page
- Go to ***Pages > Management***
- In the ***Service*** row, click on the Menu/Kebab button (3 dots) and select ***DESIGN***
- From the right column, drag and drop the new widget into an open frame in the page
- Select ***PREVIEW*** at the top of the page

You should now be able to see "Hello World" on the page!

---

## 5. Publish the updated page
- Go to ***Pages > Management***
  - Note that the status of the Services page is now yellow*
- Click on the Menu/Kebab button (3 dots) and select ***PUBLISH***

---