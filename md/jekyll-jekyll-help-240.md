# [condition.rb:18:in &#x60;block in &lt;class:Condition&gt;&#x27;: undefined method &#x60;include?&#x27; for ??? (NoMethodError)](https://github.com/jekyll/jekyll-help/issues/240)

> state: **open** opened by: **** on: ****

Liquid Exception: undefined method &#x60;include?&#x27; for 1111884:Fixnum in _layouts/bug.html
/var/lib/gems/1.9.1/gems/liquid-2.6.1/lib/liquid/condition.rb:18:in &#x60;block in &lt;class:Condition&gt;&#x27;: undefined method &#x60;include?&#x27; for 1111884:Fixnum (NoMethodError)

Beginner to Ruby and Jekyll here.  Not sure if this is a real bug of my fault.  Maybe a dependency is missing?  I have a layout (bug.html) that uses the Liquid &quot;contains&quot; operator and it fails there.  If I use any other operator, i.e. == or != for instance, it works fine.

If a dependency is missing, Jekyll should complain about it instead of launching.  Jekyll 2.5.3 used.  Version of Gems and Liquid are 1.9.1 and 2.6.1 respectively, according to the path given in the error message.  I use Ubuntu 14.04 LTS Trusty Thar.

### Comments

---
> from: [**kansaichris**](https://github.com/jekyll/jekyll-help/issues/240#issuecomment-70960509) on: **Wednesday, January 21, 2015**

Could you paste the content of your layout (&#x60;bug.html&#x60;)?
---
> from: [**deragon**](https://github.com/jekyll/jekyll-help/issues/240#issuecomment-71373775) on: **Sunday, January 25, 2015**

Sorry for the late reply, your question found its self in gmail&#x27;s spam box.  Here is the content.

    ---
    layout: default
    ---
    &lt;div class=&quot;post&quot;&gt;

      &lt;header class=&quot;post-header&quot;&gt;
        &lt;p class=&quot;post-meta&quot;&gt;{{ page.date | date: &quot;%b %-d, %Y&quot; }}{% if page.author %} • {{ page.author }}{% endif %}{% if page.meta %} • {{ page.meta }}{% endif %}&lt;/p&gt;

        &lt;table class=&quot;bug-table&quot;&gt;
          {% for data in page.bug %}
          &lt;tr class=&quot;bug-table&quot;&gt;
            &lt;td&gt;
              {{ data[0] }}
            &lt;/td&gt;
            &lt;td&gt;
              {% if data[1] != nil %}
                {% if data[1] contains &quot;http&quot; %}
                  Bonjourldskf
                {% endif %}
              {% endif %}
              {{ data[1] }}
            &lt;/td&gt;
          &lt;/tr&gt;
          {% endfor %}

          &lt;tr class=&quot;bug-table&quot;&gt;
            &lt;td&gt;
            Manual Title
            &lt;/td&gt;
            &lt;td&gt;
            {{ page.bug.title }}
            &lt;/td&gt;
              {% if data[1] contains &#x27;http&#x27; %}
              a href=&quot;{{ data[1] }}&quot;
              {% endif %}
          &lt;/tr&gt;
          &lt;tr class=&quot;bug-table&quot;&gt;
            &lt;td&gt;
              Number
            &lt;/td&gt;
            &lt;td&gt;
              {{ page.bug.number }}
            &lt;/td&gt;
          &lt;/tr&gt;
        &lt;/table&gt;

        &lt;p class=&quot;bug-meta&quot;&gt;{{ page.launchpad }} &lt;/p&gt;
        &lt;p class=&quot;bug-meta&quot;&gt;{{ page.upstream}} &lt;/p&gt;
        &lt;p class=&quot;bug-meta&quot;&gt;page.bug.title = {{ page.bug.title }} &lt;/p&gt;
        &lt;p class=&quot;bug-meta&quot;&gt;page.bug.number = {{ page.bug.number }} &lt;/p&gt;
        &lt;p class=&quot;bug-meta&quot;&gt;{{ page.bug.system }} &lt;/p&gt;
        &lt;p class=&quot;bug-meta&quot;&gt;{{ page.bug.component }} &lt;/p&gt;
        &lt;p class=&quot;bug-meta&quot;&gt;{{ page.bug.launchpad }} &lt;/p&gt;
        &lt;p class=&quot;bug-meta&quot;&gt;{{ page.bug.upstream }} &lt;/p&gt;

      &lt;/header&gt;

      &lt;article class=&quot;post-content&quot;&gt;
        {{ content }}
      &lt;/article&gt;

    &lt;/div&gt;
