# [Jekyll not recognising file changes on &#x27;build&#x27; when run in a Vagrant VM](https://github.com/jekyll/jekyll-help/issues/151)

> state: **open** opened by: **** on: ****

As requested, I&#x27;m moving this issue here from [one I raised](https://github.com/jekyll/jekyll/issues/2923) on the main Jekyll repo.

I&#x27;m using Jekyll for the first time to create a personal blog. I do all my development in Vagrant VMs (for the consistent dev environment and clean host machine), so I&#x27;ve done the same with Jekyll.

The issue I&#x27;m having is that, fairly frequently, Jekyll doesn&#x27;t appear to recognise when files have been updated when I &#x60;jekyll build&#x60; - the &#x60;_site&#x60; directory is regenerated, but with old code. It&#x27;s like Jekyll is using some &#x27;cached&#x27; version of the source code, even though I know this isn&#x27;t the case.

This applies to all file types - .md, .html, .scss etc. I tried deleting the &#x60;.sass-cache&#x60; directory on the off-chance it would make a difference to CSS changes, but it didn&#x27;t. It can take a minute or so before changes start to be recognised. Then, after another minute or so they stop being recognised again for a while.

On the other thread @parkr said that it &quot;Sounds like your _site is being synced more slowly than jekyll is rebuilding it&quot;. I know that Vagrant syncs a directory between the VM and my host machine so this does sound plausible. He then asked &quot;Can you use something like haproxy and forwarded ports to run the server from vagrant as well as the jekyll project?&quot; As a mostly frontend developer this is already pushing my technical skills to the limit, so I&#x27;m not really sure I understand this. However, As far as I was aware the VM is already running the server and forwarding the ports to my host machine.

@troyswanson asked me to say a bit more about my set up, so here it is:

- OS X 10.9.4
- Virtual Box 4.3.12
- Vagrant 1.6.5
- [PuPHPet GUI](https://puphpet.com/) for VM config set up
- Ubuntu Trusty 14.04 LTS x64
- Forwarded ports: host &#x60;8188&#x60;, guest &#x60;22&#x60;
- File system: NFS
- Server version: Apache/2.4.10 (Ubuntu)
- VHost port: &#x60;80&#x60;
- PHP 5.6
- SSH &#x60;ForwardAgent&#x60; set to &#x60;yes&#x60; on the host machine
- Ruby 1.9.3p547
- Node 0.10.31
- Jekyll 2.3.0
- I&#x27;m also using Grunt to do a few things first (imagemin, uglify etc.) before running &#x60;jekyll build&#x60; as the last task (using [grunt-shell](https://github.com/sindresorhus/grunt-shell))

I&#x27;m not sure what else to tell you. If there&#x27;s anything else it would help to know please do let me know.

Sods law, at the moment I can&#x27;t reproduce the issue. But you can be sure the next time I sit down to do a bit of dev it&#x27;ll reappear again ;)

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/151#issuecomment-55662122) on: **Monday, September 15, 2014**

Would it be possible to share your &#x60;Vagrantfile&#x60; so that I can try to reproduce this?
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/151#issuecomment-55662534) on: **Monday, September 15, 2014**

Also, it seems there might be a bug related to this specific use case:

From the [Vagrant Docs](https://docs.vagrantup.com/v2/synced-folders/virtualbox.html):

&gt; There is a [VirtualBox bug](https://github.com/mitchellh/vagrant/issues/351#issuecomment-1339640) related to &#x60;sendfile&#x60; which can result in corrupted or non-updating files. You should deactivate &#x60;sendfile&#x60; in any web servers you may be running.

Not sure if this is related, just documenting it here for now in case it is.
---
> from: [**matt-bailey**](https://github.com/jekyll/jekyll-help/issues/151#issuecomment-55662657) on: **Monday, September 15, 2014**

That&#x27;s very kind of you, thank you. I can&#x27;t do it now as I have to shut my Mac down, but I&#x27;ll send the Vagrant config stuff through tomorrow.

Interesting about that bug - might be worth looking into...
---
> from: [**matt-bailey**](https://github.com/jekyll/jekyll-help/issues/151#issuecomment-56111458) on: **Thursday, September 18, 2014**

Hi @troyswanson, I&#x27;ve zipped up [my Vagrant config here](http://cl.ly/2u1p1n1R2P1w/download/test-vagrant-config.zip).

A few points that might help:

- The config was originally generated using the [PuPHPet](https://puphpet.com/) GUI
- The VM settings can be found in &#x60;site/puphpet/config.yaml&#x60;
- The main frontend dependencies are installed using &#x60;site/puphpet/files/startup-once/frontend-tools.sh&#x60;
- Although not really necessary on this project, the shared directory is set to one level above the site root
- I&#x27;ve set the domain to &#x60;my-project.dev&#x60; and the synced target folder name to &#x60;my-project&#x60; (you can of course change all this in &#x60;site/puphpet/config.yaml&#x60;)

FYI, I did a bit more work on my site this evening and the problem was still persisting - I would make a change to one of the &#x60;.scss&#x60; files and it wouldn&#x27;t be recognised when running &#x60;jekyll build&#x60;.
