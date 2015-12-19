# [Assigning multiple collections to a filtered loop variable](https://github.com/jekyll/jekyll-help/issues/238)

> state: **open** opened by: **** on: ****

I have a couple collections set up, &#x27;case-studies&#x27; and &#x27;projects&#x27;, and I have front matter variables within them for things like client name, project date, project name, etc. I have successfully created loops that display only projects from a specific client, sorted by date, for example, but would I would like to do is create a loop that displays both case studies and projects from a specific client, sorted by date. I&#x27;m curious if there is a way to assign more than one collection to a variable to be passed to the loop? Something like the following:

&lt;pre&gt;{% assign client_timeline = site.projects and site.case-studies | where: &#x27;client&#x27;, page.client | sort: &#x27;date&#x27; %}

{% for timeline_item in client_timeline %}
{% endfor %}&lt;/pre&gt;

Everything here has proven successful except for the second collection, which is ignored.

Assigning filtered variables to loops seems like a pretty powerful feature, but I haven&#x27;t found a lot of documentation on what all can be done with them. I&#x27;ve just started with Jekyll last week so any resources with advanced usage documentation would be greatly appreciated!

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/238#issuecomment-69824018) on: **Tuesday, January 13, 2015**

You&#x27;re getting into some pretty advanced Jekyll stuff pretty quick! Bravo.

I think this (lack of) behavior is being blocked by https://github.com/Shopify/liquid/issues/427.
---
> from: [**graemeduckett**](https://github.com/jekyll/jekyll-help/issues/238#issuecomment-69844763) on: **Tuesday, January 13, 2015**

Thanks for your reply. That is disheartening news! I was feeling so close. I see there hasn&#x27;t been much action on the topic since August but here&#x27;s hoping it gets implemented.
