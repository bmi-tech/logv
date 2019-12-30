# Logv 文档

## 概述

logv is logger of vdms project for golang.

logv didn't impelement its own interface, it neseted logrus.Logger, so all logrus logger interface 
is avaliable. For example:
Warn() Debugln() Errorf()

## 使用方法

### 方法一

    调用logv.New()， 创建logv.Logger实例，使用该实例

### 方法二

    调用logv.Errorf()或logv.Debugf()等接口（推荐此方法）,使用SetLogger函数可以设置全局logger属性

    注意：
    1. SetRotate函数必须在SetOutputFile之后调用
    2. SetOutput会移除默认的日志回滚
