#!/usr/bin/python
# -*- coding: utf-8 -*-

import os,re,sys,string
import dircache
import re,xml.dom.minidom,codecs

RESOURCE_FILE = "resource/resource.json"
INDEX_FILE = "index.json"
OUT_FILE = "out/"

if __name__ == "__main__":
	ret = os.system('git pull')
	ret = os.system('git commit . -m "make version"')
	ret = os.system('git push')
	ret = os.popen('git log --name-only --pretty=format:"%ad" --date=raw')
	ret = ret.read()
	ret = ret.replace(' +0800','')
	commits = ret.split('\n')
	time = 0
	fdict = {}
	for commit in commits:
		if time == 0:
			time = commit
		else:
			if commit == "":
				time = 0
			elif fdict.has_key(commit) == False:
				fdict[commit]=time
				print commit+"."+time
				os.system("cp -vf %s %s/%s.%s"%(commit,OUT_FILE,commit,time))
	for k , v in fdict.iteritems():
		print k,v

