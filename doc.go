// Package fxtabs collects open tabs on all Mozilla Firefox windows.
//
// Tabs are collected from "recovery.jsonlz4" file, where Firefox uses as a persistent
// backup of open tabs, back and forward button pages, cookies, forms, and other session
// data.
//
// This file os written almost in real time (there will be only some seconds delay).
package fxtabs
