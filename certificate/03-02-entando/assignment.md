---
slug: 02-entando
id: 5ahjkyor5p2j
type: challenge
title: 'Welcome Wizard: Creating a Page'
teaser: 'Welcome Wizard: Creating a Page containing content and other Entando components'
notes:
- type: text
  contents: |-
    # Page Management
    The App Builder provides the capability to publish application **pages** containing content as well as other Entando components.

    **Page Templates** are used to layout the available frames on a page
tabs:
- title: App Builder
  type: website
  hostname: kubernetes-vm
  url: https://kubernetes-vm.${_SANDBOX_ID}.instruqt.io/app-builder/
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 1200
---
In the Left Sidebar, Click on *Pages > Management*

Here, you'll see a list of pages that represent the page tree of your site.
Click Add in the bottom-right to create a new page.

## Enter the following required details:

- **Title:** used for SEO

- **Code:** must be unique

- **Page placement:** location of the Page in the Page Tree

- **Owner Group:** Group that has access to the Page

- **Page Template:** structure & presentation of the Page


Click on the ***Save and Design*** button at the bottom.

Once we save our settings, we can start designing our page with widgets and other components!

---

# Now we are in the Designer Section.

## Adding Widgets:

In the Right Sidebar, Expand the Page menu
Drag & drop the Logo widget into the dotted grey 'Logo' section on the page.

Click Next to add pre-configured Navigation Menu widget to the page design for top level horizontal menu

In the Right Sidebar, expand the CMS menu
Drag & drop the Search Form widget into the dotted grey 'Search Form' section on the page.

In the Right Sidebar, expand the System menu
Drag & drop the Login widget into the dotted grey 'Login' section on the page.

Click Next to add pre-configured Content widget to the page
Click Next to add pre-configured Content List widget to the page
Click Next to add pre-configured Navigation Menu widget to add link for sitemap in the footer
---

Click ***Preview*** at the top to see what your page will look like before you publish!

Click ***Publish*** at the bottom to see your page live.

---