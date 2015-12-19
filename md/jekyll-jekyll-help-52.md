# [How to get a document&#x27;s (page&#x27;s) collection ](https://github.com/jekyll/jekyll-help/issues/52)

> state: **closed** opened by: **** on: ****

I&#x27;m not seeing anything in the docs or code on this. 

I&#x27;d like to do something like...&#x60;&#x60;&#x60; {{page.collection_label}}  or {{page.collection_name}}&#x60;&#x60;&#x60; to output the collection of the current page&#x27;s document. 

Essentially, I want to accomplish two things: 1) Put a label for the collection at the top of the page and 2) create a menu of other items in the collection.



### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/52#issuecomment-43955034) on: **Thursday, May 22, 2014**

We&#x27;d add this to &#x60;Document#to_liquid&#x60;. I&#x27;m cool with this addition.
