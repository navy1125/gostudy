#!/usr/bin/python
# -*- coding: utf-8 -*-

import os,re,sys,string
import dircache
import re,xml.dom.minidom,codecs

RESOURCE_FILE = "resource/resource.json"
INDEX_FILE = "index.json"
OUT_FILE = "out/"

def changefile(dstdir,dstfile,dic):
	global changenum
	global changefilenum
	print(dstfile)
	dst = open(dstdir + "/" + dstfile,"r")
	dstlines = dst.readlines()
	dst.close()
	dst = open(dstdir + "/" + dstfile,"w")
	i = 0
	j = 0
	first = True
	for dstline in dstlines:
		tmp = dstline
		for k , v in dic.iteritems():
			tmp = tmp.replace(k,"%s?v=%u"%(k , v))
			if tmp != dstline:
				dstline = tmp
		dst.write(dstline)

	dst.close()

if __name__ == "__main__":
	ret = os.system('git pull')
	ret = os.system('git commit . -m "make version"')
	ret = os.system('git push')
	ret = os.popen('git log --name-only --pretty=format:"%ad" --date=raw')
	ret = ret.read()
	ret = ret.replace(' +0800','')
	commits = ret.split('\n')
	time = 0
	rssdict = {}
	jsdict = {}
	jsondict = {}
	for commit in commits:
		if time == 0:
			time = commit
		else:
			if commit == "":
				time = 0
			elif re.search(".js$",file) != None:
				if jsict.has_key(commit) == False:
					jsict[commit]=time
					print commit+"."+time
			elif re.search(".json$",file) != None:
				if jsonict.has_key(commit) == False:
					jsonict[commit]=time
					print commit+"."+time
			else:
				if rsonict.has_key(commit) == False:
					rsonict[commit]=time
					print commit+"."+time
	changefile("./","index.html",jsdict)

