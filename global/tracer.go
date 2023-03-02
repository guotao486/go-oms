/*
 * @Author: GG
 * @Date: 2023-02-19 10:54:25
 * @LastEditTime: 2023-02-19 10:55:04
 * @LastEditors: GG
 * @Description: Tracer 全局对象
 * @FilePath: \oms\global\tracer.go
 *
 */
package global

import opentracing "github.com/opentracing/opentracing-go"

var (
	Tracer opentracing.Tracer
)
