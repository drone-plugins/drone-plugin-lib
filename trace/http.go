// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package trace

import (
	"context"
	"crypto/tls"
	"net/http/httptrace"
	"net/textproto"

	"github.com/sirupsen/logrus"
)

// HTTP uses httptrace to log all network activity for HTTP requests.
func HTTP(ctx context.Context) context.Context {
	return httptrace.WithClientTrace(ctx, &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			logrus.WithField("host-port", hostPort).Trace("ClientTrace.GetConn")
		},

		GotConn: func(connInfo httptrace.GotConnInfo) {
			logrus.WithFields(logrus.Fields{
				"local-address":  connInfo.Conn.LocalAddr(),
				"remote-address": connInfo.Conn.RemoteAddr(),
				"reused":         connInfo.Reused,
				"was-idle":       connInfo.WasIdle,
				"idle-time":      connInfo.IdleTime,
			}).Trace("ClientTrace.GoConn")
		},

		PutIdleConn: func(err error) {
			logrus.WithField("error", err).Trace("ClientTrace.GoConn")
		},

		GotFirstResponseByte: func() {
			logrus.Trace("ClientTrace.GotFirstResponseByte")
		},

		Got100Continue: func() {
			logrus.Trace("ClientTrace.Got100Continue")
		},

		Got1xxResponse: func(code int, header textproto.MIMEHeader) error {
			logrus.WithFields(logrus.Fields{
				"code":   code,
				"header": header,
			}).Trace("ClientTrace.Got1xxxResponse")
			return nil
		},

		DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
			logrus.WithField("host", dnsInfo.Host).Trace("ClientTrace.DNSStart")
		},

		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			logrus.WithFields(logrus.Fields{
				"addresses": dnsInfo.Addrs,
				"error":     dnsInfo.Err,
				"coalesced": dnsInfo.Coalesced,
			}).Trace("ClientTrace.DNSDone")
		},

		ConnectStart: func(network, addr string) {
			logrus.WithFields(logrus.Fields{
				"network": network,
				"address": addr,
			}).Trace("ClientTrace.ConnectStart")
		},

		ConnectDone: func(network, addr string, err error) {
			logrus.WithFields(logrus.Fields{
				"network": network,
				"address": addr,
				"error":   err,
			}).Trace("ClientTrace.ConnectDone")
		},

		TLSHandshakeStart: func() {
			logrus.Trace("ClientTrace.TLSHandshakeStart")
		},

		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			logrus.WithFields(logrus.Fields{
				"version":                       cs.Version,
				"handshake-complete":            cs.HandshakeComplete,
				"did-resume":                    cs.DidResume,
				"cipher-suite":                  cs.CipherSuite,
				"negotiated-protocol":           cs.NegotiatedProtocol,
				"negotiated-protocol-is-mutual": cs.NegotiatedProtocolIsMutual,
				"server-name":                   cs.ServerName,
				"error":                         err,
			}).Trace("ClientTrace.TLSHandshakeDone")
		},

		WroteHeaderField: func(key string, value []string) {
			logrus.WithFields(logrus.Fields{
				"key":    key,
				"values": value,
			}).Trace("ClientTrace.WroteHeaderField")
		},

		WroteHeaders: func() {
			logrus.Trace("ClientTrace.WroteHeaders")
		},

		Wait100Continue: func() {
			logrus.Trace("ClientTrace.Wait100Continue")
		},

		WroteRequest: func(reqInfo httptrace.WroteRequestInfo) {
			logrus.WithField("error", reqInfo.Err).Trace("ClientTrace.WroteRequest")
		},
	})
}
