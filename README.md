# 说明

logv is logger of vdms project for golang.

logv didn't impelement its own interface, it neseted logrus.Logger, so all logrus logger interface 
is avaliable. For example:
Warn() Debugln() Errorf()
