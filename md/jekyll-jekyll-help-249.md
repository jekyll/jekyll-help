# [Re-usability of posts (and include of external posts) ?](https://github.com/jekyll/jekyll-help/issues/249)

> state: **open** opened by: **** on: ****

Hi, 

So one of the things I would like to have is  a way to make a &#x27;_posts&#x27; available to other modules without duplication.  But also &#x27;import&#x27; posts.  For context, we&#x27;re building education &amp; training modules, where &#x27;posts&#x27; are training components relating to the module.  We have more than one instance of Jeykll and want to &#x27;share content&#x27;.

Currently: 
1) One post called &#x27;conflict resolution&#x27; in training module 1, is also used by two other modules, and so it must be copied to each of those modlues.
2) If content for that post needs to be  updated, I have to copy it across all three instances.
3) if someone wants to contribute a change to that post ,   they have to submit a PR to all 3, or I have to copy their PR to all three.
4) Additionally, we are running more than one instance of Jeykll and may want to share a post, so would like to &#x27;include&#x27; an external post ie : https://raw.githubusercontent.com/emmairwin/reps-edu-cirriculum/master/modules/conflict_resolution.markdown    vrs make yet another (now external) copy.


So this is about modularity  and maintainability of content, but to encourage quality of content through participation (contribution). I know that&#x27;s not a &#x27;blog thing&#x27;, but it is a community thing, and a way that invites contribution and re-use for community projects.  

I&#x27;m new to Ruby(but long-time web-dev), so have been playing with this but not quite sure what&#x27;s possible. I&#x27;m sure there are holes in what I am proposing - but I really need a more modular/reusable way of working with posts. I am playing with the code, but if there are suggestions that would be helpful


### Comments

---
> from: [**emmairwin**](https://github.com/jekyll/jekyll-help/issues/249#issuecomment-71387791) on: **Sunday, January 25, 2015**

So currently working with this solution (based on include function)

I  created a stand-alone repository for common_content  , so community can work on content without worrying about the actual implications to jekyll.  This way I can also tag releases, and have branches related to content.

I clone common_content into the _includes folder and then simply use  {% include common_content/test.markdown %}   .   

I think this should work.  Although the whole branching tagging bit is probably overthinking it a bit.  



