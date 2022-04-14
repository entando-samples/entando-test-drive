---
slug: 03-entando
id: dmzxjtxmxkxs
type: challenge
title: 'S&P Returns Dashboard: Create the Page'
teaser: In the previous challenge, we installed  the technical solution to compose
  the Dashboard Page asked by our Consumers, a Tableau PBC. This PBC allows us to
  display the requested Financial information using a Micro Frontend component.
notes:
- type: text
  contents: |
  **CREATE A PAGE FOR THE PBC**
  
  Now that the Tableau PBC is available in the App builder, the Composer can use it to build their application. In this challenge, we will create a new page in our application and prepare it for adding the PBC.
tabs:
- title: App Builder
  type: website
  path: /app-builder/
  url: https://named-k3s.${_SANDBOX_ID}.instruqt.io/app-builder/
- title: shell
  type: terminal
  hostname: named-k3s
difficulty: basic
timelimit: 1200
---
**CREATE A PAGE FOR THE PBC**

Now that the Tableau PBC is available, we can add it to a page. This PBC contains a micro frontend that displays the S&P Returns and Economic Indicators.

In the Left Sidebar, **Click: Pages > Management**

Here you will see the current page tree for your application. **<span style="text-decoration:underline;">Click: Add </span>**in the bottom-right to create a new page.

**<span style="text-decoration:underline;">Enter the following required details:</span>**

  - In Title: **<span style="text-decoration:underline;">S&P Tableau</span>**. 

  - In the page tree section, **<span style="text-decoration:underline;">select the page emplacement by clicking on Home</span>**, (it should be highlighted in blue).

  - In Page Groups: **<span style="text-decoration:underline;">change Owner group from Administrator to Free Access</span>**

  - In Settings **<span style="text-decoration:underline;">choose 1-2 Columns Page Template</span>**.

Then**<span style="text-decoration:underline;"> click: Save and Design</span>** button at the bottom, (the Designer is displayed).

Congratulations! You’ve completed the challenge. You've discovered how to create a new page and open the design mode. Next, you will discover how to drag and drop modules to the page.

**<span style="text-decoration:underline;">Click Check </span>**at the bottom right to move to the next challenge.
