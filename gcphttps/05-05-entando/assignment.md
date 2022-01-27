---
slug: 05-entando
id: lsmh1xhkny01
type: challenge
title: Creating a UX Fragment
teaser: Create a basic UX Fragment that can be reused across multiple pages
notes:
- type: text
  contents: A **UX Fragment** is a way to take a common piece of front end code and
    reuse it across multiple pages or widgets. Common elements such as basic HTML,
    javascript, or freemarker logic can be stored as fragments and referenced via
    the ```<@wp.fragment …``` tag.
tabs:
- title: Shell
  type: terminal
  hostname: entando-vm
- title: App Builder
  type: service
  hostname: entando-vm
  path: /app-builder/
  port: 80
  new_window: true
difficulty: basic
timelimit: 1200
---
## 1. Creating a new fragment
- In the App Builder, go to: ***Components > UX Fragments***
- At bottom of page, select ***ADD***
- Enter the following details:
  - **Code:** ```test```
  - **GUI Code:** ```<h2>Hello World</h2>```
- Click ***SAVE***

---

## 2. Place the fragment in a template:
- Go to: ***Pages > Templates***
- In a selected template row, click the Menu/Kebab icon and select ***EDIT***
- In the Template text box, add ```<@wp.fragment code="test"/>``` on a new line between the ```<body>``` and ```</body>``` tags
- Click ***SAVE*** to save the page template

---

## 3. View the page with the new fragment:
- Go to ***Pages > Management***
- On the ***Service*** row, on the far right side, select the Menu/Kebab button and select ***DESIGN***
- Select ***PREVIEW*** at the top of the page
> Note: You will see the fragment ```<h2> This is a fragment. </h2>``` which includes the HTML tags. By default html embedded via a fragment tag is escaped so you get it rendered exactly as you enter it. You’ll need to un-escape it to get it to render correctly.

---

## 4. Update the Fragment:
- Go To ***Pages > Page Templates***
- In a selected template row, click the Menu/Kebab icon and select ***EDIT***
- Change the tag to: ```<@wp.fragment code="test" escapeXml=false/>```

---

## 5. View the page with the updated fragment:
- Go to ***Pages > Management***
- On the ***Service*** row, select the Menu/Kebab button on the far right side and select ***DESIGN***
- Select ***PREVIEW*** at the top of the page

You should now be able to see the correctly rendered fragment!

---