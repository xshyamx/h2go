#!/bin/bash
mkdir -p h2-db
java -classpath h2.jar org.h2.tools.Server \
		 -tcp -tcpAllowOthers -ifNotExists \
		 -trace -baseDir ./h2-db
