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
    # Congratulations Composer!

    The Tableau Packaged Business Capabilities micro frontend component is now usable in our application. It's time to finalize the composition by including both the micro frontend and the content component in a page.

    Keep going. Success is right around the corner.
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
Now that the Tableau Packaged Business Capabilities has been installed, the included components are available in the platform.
This PBC contains a micro frontend that displays the S&P Returns and Economic Indicators.

In order to use it, we need to first define the page which will host it.

In the Left Sidebar, Click on **Pages** > **Management**

Here you see a list of pages that represents the page tree of your application.
Click **Add** in the bottom-right to create a new page.

Enter the following required details:

In the **Info Section** fill the **Title:** with the **S&P Tableau**.
Then, in the page tree, select the page emplacement by clicking on **Home**, it should be highlighted in blue.

In the **Page Groups** section change **Owner group** from **Administrator** to **Free Access**

Finally, in the **Settings** section choose the **1-2 Columns** **Page Template**.

Once all these fields are filled, click on the **Save and Design** button at the bottom, the Designer is displayed.

---

Congratulations! You’ve completed the challenge. You've just discovered how to create a new page and open the design mode.
Next, you will discover how to design it by drag and dropping components.

Click **Check** at the bottom right to move to next challenge.
