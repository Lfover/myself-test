package main

import (
	"fmt"
	"log"
	"os"
)

func main1() {
	// 假设你有一个Base64字符串
	base64String := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAEAAQAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//////////////v/////////+//7//v/+//7//v/9//7//v/+//////8AAAAAAAAAAAEAAQABAAEAAQABAAEAAAAAAAAAAAAAAAAAAAABAAEAAQABAAEAAQABAAIAAgADAAMAAwAEAAQABQAEAAUABQAFAAQAAwADAAIAAgABAAAAAAAAAAAA//////7//v/+//7//v/9//z//P/9//z//P/8//z//f/9//3//f/9//3//f/9//3//v/+//7//v/+////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAQABAAEAAgACAAIAAgADAAQABAAEAAQABAAEAAUABAAEAAMABAADAAMAAgACAAEAAQAAAAAAAAAAAAAAAAD///7//v/+//7//v////7///////7//v/+//7//v/+//7//v/+//7//v/+//7////////////+//7///////7/////////////////AAAAAAAAAAAAAAEAAQABAAEAAQACAAIAAgADAAMAAgADAAIAAgACAAIAAgACAAIAAgABAAEAAQABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD///////8AAP7//f/9//7//v/+//7//f/+///////+/////////////////wAAAAAAAAAAAAAAAAAA//8AAAAAAAAAAP//////////////////AAAAAAAAAAAAAAAAAAABAAEAAQABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD///////8AAAAA///+//////8AAP////8AAAEAAQAAAAEAAQABAAIAAgABAAIAAgADAAMAAwADAAMAAwADAAQAAwAEAAUABQAEAAQABQAEAAUABAAEAAMABAAEAAMAAwACAAMAAQABAAEAAAAAAP////////7//f/7//v/+//8//v/+v/6//r/+f/4//j/+P/4//f/9//4//f/9//3//f/+P/5//r/+v/7//v/+//8//3//f///wAAAAAAAAEAAwADAAMABAAFAAcABwAIAAgACQAKAAkACgALAA0ADAALAAwADQANAAwADAAMAAwADAALAAoACgAJAAkABwAHAAcABgAFAAQAAwACAAEAAAD///7/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av9f/0v/Qf9B/07/ZP9P/z//PP9D/1f/bv9l/2H/b/9w/4z/mf+i/67/uv/E/9L/4//s//v/AQAWACMAKAAyAD0ARwBfAG4AbgByAHcAggCPAJ8AmgCaAJ0AnACrAKsAqwCnAKIApwCgAJ0AmwCQAIoAhAB3AHEAZgBcAFEARgA/ACwAIwAWAAcAAAD0/+H/0P/B/7b/p/+T/4X/ev99/3f/YP9P/0f/Rf9g/2H/Q/89/zf/Q/9e/2H/Wf9e/1r/bP+A/4v/n/+a/6b/s/+5/87/5f/f/+f//P8IAB0AHAAiADEARgBNAFYAWwBdAHAAewB/AIIAhwCOAJEAlQCgAKYApACgAJ4AoQCoAJ8AngCXAIsAigCDAHwAeAByAGQAWwBIAEYAQgAzACcAEwAKAPz/7//k/9T/wf+w/6z/pP+c/4v/eP9p/2f/b/9q/2D/Sv9C/03/Uv9V/1z/Tf9L/1n/Y/91/3T/d/+D/43/lP+e/6n/vP/H/8f/1v/d/+////8IABMAHwAjADMAPgBHAFsAXQBmAGgAcACDAIsAiwCVAJQAlwCgAKEArgCnAKMAnwChAKIAogCZAJIAjwCAAIQAegBsAGUAXABTAEgANAAvACQAEAAKAPz/7//e/83/x/++/6f/nv+W/4P/e/9t/27/Yv9O/1L/Tv87/03/Tf9F/03/P/9I/2D/Yv9h/3H/a/93/4T/mP+s/6D/sP/F/8z/2//z//f/AAAOABwALAA1AEAASQBaAGIAZgB0AH8AgwCMAJ8AmgCSAKkArgCvALMArACtAK8ArQCvAK0AmwCeAJkAkACIAHgAdwBrAFgATgBIAC8AKAAXAAQA///m/9n/zP+3/6n/of+S/33/Zv9l/1D/RP9c/1D/Nv8u/x//MP9Z/03/Pf8//zf/Rv9t/3z/dP9s/3z/kf+s/7v/wf/O/9b/3v/5/wwADQAnACwAMABFAEoAYgBsAGQAhgCVAIsAkACWAKwAugCzAK8AsQCzAL8AwwC7ALUArgCxAKkAnwChAJQAhwB1AGUAaABXAD8AMgAaABIABADu/+H/xf+3/6//m/+D/2b/VP9b/2T/Uf89/xz/FP86/0H/QP83/x//J/9D/0v/X/9s/1z/ZP90/4r/r//A/7T/vP/Q/+7///8IABEAHQA0ADwAPQBLAG8AcQCAAH0AeACWAKgAtACuAKMAqgDBAMIAuQC8AMMAuAC7ALAApwCqAKAAoACIAHMAbgBjAFcAQAAtACoACwDz//D/2P/I/7r/mf+B/2r/ZP+B/2T/Of8u/x3/Nf9J/z//NP8e/xP/PP9U/0//TP9N/2H/a/+B/5b/qP+l/7H/zf/g/+v/9f8TABUAHgAyAEYASQBSAHIAfAB9AH8AjACkALAAoQC0ALYAsQC/ALgAwQC9ALsAxgCrAJsAqwCnAKIAhwBtAGsAXQBUAEkAMwASAAEA8//y/97/uv+k/4r/df95/4j/Z/9D/yP/Kv9E/0n/RP8o/xf/Gf83/1f/UP9B/0v/Uf9w/4j/if+j/63/qf/R/+P/5//8/xMAIQAcADcASgBOAFcAcACIAI0AggCCAKcAsgCsALQAqwCoALUAugDFALAApgCzAKYAmwChAJoAjABxAGEAaABPAD8AMwAeAAUA9f/r/97/wP+g/47/gP+G/4X/af89/yz/Nf9e/1H/LP8q/yH/Nf9O/0z/Rv9Q/07/bf98/3j/o/+m/6P/w//U//P/9v/0/xsAIwA1AD8ATABaAGMAdgCQAIwAhQCdAJ8AsACyALIAswCpAK0AwgC0AKwAswCbAKAAlgCNAJEAcgBlAGAAQgA7ADEAHwAMAO3/4//b/77/o/+V/4D/gP+G/2j/RP81/zr/WP9V/z3/Mv8m/y7/VP9i/1D/Tf9R/3n/f/+K/6P/sP+8/7v/1f/4/wAACgAgABwANwBHAFkAYABfAH0AjACPAIcAmQCkAKgAqQCuAK0AowCrALEAqgCbAJwApACNAHsAfgB5AGcAUgBFADwAKQASAAwA9//q/9D/u/+l/4n/j/+P/5T/XP8//0b/RP9d/2n/TP8v/zP/Qf9k/2T/Wv9g/2//e/+L/5v/rP+6/8f/3f/h//D/CAAlABkALABLAE8AWQBcAHEAgwCHAIwAmACVAJcAmwCrAKkAnACqAJkAkACZAJcAlACPAHYAbwBoAFgAWABHADoAIQANAAsA+//p/9b/uv+l/5b/hP+R/5L/bv9g/0L/RP9t/2T/X/9R/zP/VP9r/23/cf9n/3v/if+Y/67/rv/A/9T/1P/4//7/CgAiACAALABDAFsAYQBjAGYAdgCEAJEAkwCTAJMAkgCcAKIAmwCRAJIAjACPAIcAgAB4AG8AZABdAFAAPgA6ACUAFwAKAP//7P/Z/8X/tP+i/5T/eP+J/5P/b/9e/z7/Rf9i/3P/bf9W/0v/V/9j/4L/gv90/5z/k/+n/7X/uv/d/+L/6v/9/wsAIAAqADMAPgBFAG4AawBqAHwAeQCFAJIAlAChAJsAiQCTAJgAlQCQAJAAgQB2AHQAeABpAF8ATgA9ADwAJQAhABgA/P/o/+b/zv/F/7P/mP+D/2v/fv+L/4r/V/82/0v/Wf93/4T/XP9T/2L/bf+T/4n/jv+g/53/u/+//9X/7P/f//v/EAAZADQANQA0AEYAVwB1AHwAcAB9AHgAiQCXAJYAqgCQAIUAkQCQAJAAjACBAHcAcQBnAGIAXQBLADsANAAgABoADQD9/+n/1v/O/8H/tP+e/4v/cf9h/3D/l/95/0//RP9A/23/fv92/2//aP9u/4//jf+U/6v/rv/H/8j/0f/x//r//f8TABwANABBAEgASgBKAHIAhgCCAHoAfwB/AIwAjgCWAJ4AhwCGAIIAfgCBAHYAcABnAFkAYgBTAD8ANgAeABUAFAABAPn/4v/H/8H/uP+r/5z/kf9v/1r/Wv+J/4v/Z/9S/zr/ZP+D/4r/f/97/3//mv+j/57/s//B/9v/2v/g//v/CQALAB8AIAA/AFgAWABcAE0AdQCEAIgAjgB8AIoAkwCGAJYAjgCIAI0AegB9AHAAdABpAE4AVABNAEAARAAeAAoABADv//f/5P/Q/7r/rf+j/5n/jv95/17/WP9f/4D/kP9a/1P/Sv9s/5f/jf+R/5L/kf+o/7r/s//L/9T/6v/3//X/FgAXABsAJwAxAFsAXwBgAGoAYgB1AH8AjQCNAIoAhQCHAIEAgwCMAH0AfQBnAGsAagBZAFMASAA3ADIAKwAlAA4A8v/v/9v/1//G/7z/rP+a/5D/jP95/2L/T/9h/43/ef9+/17/WP9+/4//oP+h/5v/q//A/7T/yv/U/+T/9f/8/w4AGQAgACMALgA8AFAAZgBsAF0AagB2AIEAggB7AIcAgAB9AIAAdAB1AHAAaABmAFQAWwBVAD4ALgAtACoAFwAKAP3/7P/g/9P/v/+5/63/pf+b/4//fP92/1L/SP+D/43/jP9u/1f/ff+g/6f/sv+d/6b/yv/M/9v/zv/j//v//P8PACMAJAAlACkAOQBdAGEAbQBiAFwAeACGAIYAfwBxAHwAiAB/AHkAbgBgAFsAZQBeAFMARAA1ACYAIwAdABQAAwDt/+X/2P/S/7j/sv+n/5n/mf+U/37/a/9F/03/lv+U/5L/ef9t/4n/pP+w/7z/uf/F/+f/3P/l/+j/+v8DAAMAKAA+ADMAOAAvADgAXABZAHYAYwBgAHsAcAB/AHgAcQB2AGsAdgB4AGYAXQBMAFIASQA8ADUAGwAaABIACAADAOv/4f/T/73/uf+6/7D/nv+W/5H/if9q/0n/Xv+O/5f/kv+D/3H/lP+2/7H/tf/B/9X/7v/i//D//P/2/wMADwAiADQAQQBAADEANgBUAF0AVQBWAGkAdwBvAHEAcQBoAG4AdQBnAGQAZgBdAFIARQA4ADEAHwAaABoA///x/+3/6f/V/83/wP+v/7L/rP+e/5v/mP+A/23/Xf+D/57/lP+J/4b/n/+8/8r/uv+//87/7//+//7/CQAIAAQADgAmAC8ANgA/AEMAPgBGAEkAUQBTAFkAZgBWAF8AdAB4AGgAWwBWAF8AYQBdAEwAPQBFADgAIgD+//X//f/6//b/6//h/9b/xf+u/6j/mv+W/6j/oP+R/33/dP+N/5H/kv+P/5H/sP/K/9v/zf/G/9z/5//2/wYABQAVABwAHwAoACUANgA/ADoAMgA7AE0AVQBOAEgATwBSAFwAWgBXAF4AYwBgAFQAUQBPAFQATQAzACUAIgAiABQAAADz/+L/4f/g/87/wv+5/7j/tv+p/6P/n/+Y/4v/hv9v/3j/m/+X/5v/mP+j/7//yv/K/9f/4v/t/w8ABgAQACwAHwAlACwAKQA+AE0ARwBDAD4ARQBKAEgAQAA5AEYAUABaAFIAWgBoAFcAUgBWAFsAVgA+ACsAIwAgABoA+//e/9f/3P/f/9n/zP+//8H/r/+y/7b/o/+j/5L/hf+L/5v/nv+V/4z/h/+w/8j/0P/S/8r/5v/y/wsADwAHABQAGwAtAC0APQBDADsAOQA8AEEAQABBADkAPAA6AEcARwBDAE0ASgBaAFUAVgBYAEgASwBLADwAIQAXAAUA+P/v/+H/2//X/9H/zP/P/8n/0P/B/7H/sv+s/53/kP+N/5D/qf+c/5v/qv+u/8r/0v/Q/9j/8v/5/wgAHAAVABkAHAAkADQAMQAuADwAMAAzADwANgA3AC8AJgAuAEUASABMAEUAQQBLAEYASABBAD0ARABGADYAKwAbAAYA/P/i/9b/z//J/8r/yP/B/7//xf/D/7//sf+s/6//s/+z/6D/of+0/8D/v/+5/7f/0v/m/+X/7P/u/wUAFwAaABYAHAAtADwARQAzADcAPQA8AD0AMgA0ADcAMwAsACcALwA5ADoALwAxAEEASQBRAEcANAAyAC8AIgATAPr/7//x/+P/5P/a/9L/0//P/9D/0v/V/8b/vv+1/7D/vf+p/6D/nf+X/7b/wv/N/9P/yv/d/+n/8v8DAAMABQAPAB0APQA/ADMAJQAnADYAOgA/ACsAIgAgACgAMAAsACwAJAAqADgAOwA+ADsANQA6AD0APgArABkABgD7//X/2P/a/93/5P/i/9b/3//g/93/1//U/9X/2//H/7L/sv+1/8L/uv+u/6f/uf/Y/9//6v/d/+j/AwAQABYABgAXABoAHgAjACUAMwAwADcAJQAjACkAKAAnABMAIwAuACoAKAAiACwAOwA+ADoALAAoACoAJQAUAP//9v/r/+z/6//l/+j/7//t/+//8P/l/+H/1//R/8z/zP/O/8//x//C/9D/0f/O/9L/yP/R/+7/9//5/+z/AwAbACIANAAaABoAMAA6ADYAFQACAPz/EAAXAAoAAAAAACAAJwAfACcAKQA2ADYAGgALABcAIAAVAPb/6//4//r/+f/x/+j/9v8LAAYA8//n//z/AwD3/+v/2v/i/+j/2v/J/8b/yv/T/8//0P/W/+b//P8EABAACQAYACoALgAlABgAFwAUAA8A///4//r/AgAAAAQAEAAYAB0AFQAZACEAIgAhACsAIQAVAB4AEwAHAAUAAgACAPX/7P/y/+z/7v/s/+z/9P/v//T/9////wAA+P/r/9v/2f/c/9//1v/P/+r/BwAQABMAGgAkACEACwADAA0ACQD6//H/6P/2/wcA8f/l//z/BAAGABMAGQAgACMAIwAiAB0ADQD6//b/+/8IAAMAAQAAAAgACQDv/+v/4f/6/xIA+f/r/+//AAALAAgA9v/u/+//+P/5/9v/2P/o/wcAAwD1/wgADwAVAAMA+f/q/+n/9P/0//7/+/8YABUAAwAGAAUACQAAAP3/AQAPABMACwAPAB8AHwASAA8AAgAAAAEA+P/9//z/9v/1//3/AwD3/+//6//1//3/7f/2//3/AgAVAB8AJAAXABYAEAD+/+P/2//i/+D/4//u/wQAIQA3ACMAAgDw//3/+v/d/8r/v//c/wAACAAHAAQAEQAuADsAKwASAAAA+P/9//f/7v/5/wsAIQAfABAADwADAPL/6P/n/9//5v/u/+T/+/8RABsAEQANABAACQAFAPD/6//r//b/AgAJABEAEQAWABYADgD8/+3/8f///wAA/v8DAA4AEQAJAAIA/P/x/97/zP/B/8r/1f/i/wEAKQBJAFMATwBEADUAKQARAPT/5P/f/+z//f/y/9//2v/W/8j/vf+4/8z/6P/i////IwA3AEIANAAuACYAJAAJAOL/3//j/93/1P/e/+//+v8JABAAGQAZAP//6P/y/wIAAgAEAP3/AwAkADoANgARAPn/9v/4/+X/zv/J/8//8P8EAAMADAAvAEcAQgA/ADUAIQATAA0A+//j/9n/3P/e/+L/6v/t/+P/2v/z/woA/P/o/+r/AAAOABIAEwASABUAEQAOAAEA6//f/+z///8IAA0AEAAkACMADAAHAAAA9v/x/+f/2f/s////AgAIAPX/8v8HAAYA9//3//v/9v/7/wMAEgAfAB4AJAAhACAAIgAHAPb/6//d/9X/3f/8/wIACwAZACIANwAyAA8A8f/u/+P/3//j/93/4//1/wQA+//7//3/BwACAO////8HAAAA/P8GAAwAGgAeAAQA/P/+/wUAAQDt/+P/6//z//D/6v/3/wwAEgAGAOz/9f8EAOr/3P/o//f/9P/5/woADAALAAYAAwAIAAIABwADAP3/EQAKAAYADAAPAAwAEAAJAPb/AAAAAAAA/v/1//z/FgAcABcADwD+//z/6f/T/9z/7P/1/wIADwAjACoAJQAZAAoA+//y//P/4v/d//P/FQAVAA8AHgAnAAwA+f8JAP3/6//l/+L/6v/u//r/AwABAAwAHAAfAAsA/P8AAPb/5//e/+n/DwAaAP3/9v8SABYADgD//+n//P8FAOr/2v/w//f/CAASAPz/DQAkABYAEgARAP//BAAFAOz/1//W/+f/8//5/wAABAARABoAFwAeAA8ABwAEAPH/8/8DAPP/7P8EAAgAAADw//L/BgD//97/6P8EAAYAEQAKAAEADAAbABoAAgDl/+X/9f/w/+P/7v8AAPn/DgAVABQAIwAPAAUADwAYAAEA6P/j/+X/5v/f/+T/8f8JABQADgAMAA8AIQAmABgAFAAPAAQA9P/T/9P/1//k//H/4f/v/xAAFAAAAAgABQAOABQACQANAP7//P8EAAEA7v/w//r/+f/w//X/AQABAA0ABgAPABYA/////wAA7P/n//H/9v8JAAMABQASABgAJgAWAAAA9//1//r/9P/n//L/AAANAAkACwAiABIAEAAIAPL//P8IAAAA6f/l//z/AQD6/wAACgAKAPn//v8GAA0AFAAFAPX/9/8TAP7/6P/v//b/CwD6//n/CQADAAkAEgAPAAAA8f8KABUAAgDp/+T/AgAIAA8ABQD9/xYAIQALAPz/4P/L/+n//f8DAPz/DAAZAB0AKAAaAAQA+v/5//X/9f/s/wkAHAABAAQADQAIAPT/9/8AAOL/4//t/+X/9f//////CAARAB8AJAAZAAsAEQASAP7/8P/z//7/7v/m/+r/6v/5/wMA/v/z/woAJgAfAA0ADQAGAAAA9//n/+7/8v/z//j/AwABAA0AEgD9/wIAAAD3//7/+f/z/wAAAgAHAAQAAAAEAAAA/f/0/+z/7//4//L/9v/6//X/EQAWAP3/AwATABQADAAJAP3/7v/5/wEABwD0/+H/5//x//z/9//9/wcACgARABMAGwAIAP7/9//t//r/8P/y//v/+v/0/wAAAAAAAAsAAQAPAP//+P8KAAoACwD+/wAABQABAAQA+v/2//7/8f/z//L/9v8NABIAFQAJABQAJwAeAAoA9v/w//b/8f/j/9v/5P8FAA0ACQAEAAoAJAAhAAAA+/8HAPv/9v/u//H/5v/w/wEA/P8KAAsABQAHABIAHwAOAPH/+f8HAAAA9P/n/+//9/8NAAwA/P8EAAkAFAAKAPr/AQAJAAQA7v/u//b/6//8/wEA/v8SABUADQAMAA8AEAAIAP7//v/5/+b/6f8AAAMA+/8GAAUA//8GAAsAAADy//r//v8BAAYAAQAEAAIA/P8AABIA+//Y/+7/DgALAOz/6P/3/wUABwD7//7///8FAAwABQADAA8AFAATAA8AAwD8//X/CAACAPn//v/5/xMAEQADAO7/7P8FAPj/9v/2//f/CgAQABoABwD3/woAFAACAPT/9/////v/7v/1//z/AwAIAAAABQAaABoADQADAAQAAwDw//H/6f/g//b/+v/y//b/+/8SAA0ABAAMAA0ACwACAAMA9v/t//L/AgABAP//BwAAAAwACQD+/wQABwABAAcACwARAAoA/P/w/+///P/9//r/5f/4/w8AAAD5//v/AwACAAUAFAABAP//CwAJAAMA9f/9////AgAJAA0AEwAFAPz/CAAPAAIA/P/1/+7/8f/7/+//6v/3//z/CQADAPb/CQAGAAAACQAAAAgACwAOAAwA/P8PAAgA7P/3/+//7//6//L///8AAA0ACwD9/wIACAATAPb/8v/8//b//P/q/+z/+P8EAAQACgAXABMAKQAVAAYAFQAIAAAA7f/m/+P/4//t//T////4/+//FgAxABAAAwD9//7/BwD2/+r/5//t/wkACwABAAMAAQAUABEABwABAPn/AwAEAAAA/P8AAAcABgAGAAgAAQD3//T/8f/i/+H//f8JAAUACAAUAB4AHQARAAIA+v/9/+7/5P/l/+r/7P/t/wEABQAWABgACAAMABMAEQD2/+z/8P/8//r/8f/z/+//AwAKAAAABQAAAAMA/////wcA/f8BAPv/+/////f/AgAAAO//7f8AABMAFgAEAAgAAQANABsABgD9/97/6v8MAAIA8v/m////EgABAAUAAgD+/wsACAD///D/7f8AAAAA7/8FAAQA//8HAP//DAD+/+v/7v8CABsAAgD2/wYACwAMAAAA+//4//L/8//m//H/9f/l/+//BAALABIAFwAVABgAGQAlABAABQAFAPj//f/u/9v/3v/l//X/BwAPABYAHgAlAA0ABQAOAPn/7f/5//r/8P/v//T/9//0/wIAAgALABcADQATAAsACQATAAsA+//9//v/BAAAAPf/+//1/wAAAgAGAPv/AQD8/wwAJAANAAsABAD8//b//v/1/+L/4f/4/wYA+v/4//v/EwAUAA0AGwASAAMA+v/4//7/+P/0/+z/6/8DAA0ABQD2/+//AAADAPj//v/5//7//v8IAA4A///2//P/AgD9//j/6//0//3//P8AAP3/BAAKABMAGAAYAAYAAAAAAPX/8v/6//H/4f/z/w4ACwAIAA0ACAAcAA8ADAAJAOj/7f/7//v/8P/t//b//f8AAAsABQANABUAAwAFAAoACAAQAAoA/f/u/+//AgAAAPz/8v8AAAAA/v8GAAcACQAEABAAAwD9//L/8v8LAP7/9/////f/DQAcAP3/8//8/wQA/v/3/+//7P8AAAUABgAIAAoADwAGAAkACAD///X/6v/z/+L/8f8EAPr/EQASAAwADgANAAUA5f/r/wEA5//o//L/BgAjAAsA/P8OABYACQD0/+D/+//8//T/7v/x/woABwARAAoACQD4//f/CwD///X//v8GAA4AEAAHAAsA/f8HAAAA7P/k/+v////4/wIAAAD5/wMAEwAOAAUACAAKAAYABgABAAAA///x//H/6P/y//7/9P/7//r///8ZABYACAANAA8AAwD2//b/9P/y/+z/AAAFAO7/CAATAAUAAAD7/xMAAwDt/wIACAAAAPT/+v8AAAYAAQD0//j/6v8AABgA+f/r/wsAFQAWABEA7v/y/wUA+//o//L/9f8DAA8AAgAMAAkADQANAAEA/f/4//b/7P/9//7/8f8MABYAHgAUABAAHAAMAAIA+f/2//f/7P/i/+3/8//u////EwARABIAFgALAAUA/f8HAPX/1v/w/wUA/f/t//v/FgAdABgAAQD0/w4ACAD5/+n/4f/3//z//f/9/wAA/f/2/w4AGAAGAAYADAARAAcAAQD8//n/CAAAAPf/6//6/wQA/f/9//n/AAAQAAgABQAOAAUAAAAAAAQA/f/x/+z/8P/s/+3/7v///xIABgADAAkAFgAYABAA///1//7/AwD9//n/8P/t/+z//f8LAAAACQAOABgAGQASABUAFwAMAAAAAAACAPv/7v/w/9//6v/2//f/7f/y/wUADgAeABQADgANABAADwAMAP3/8P/6//7/+//j/+b/EAALAPn/CQANAA0ABwAFAA4ABQDv/+7//v8AAPH/6v/7//L/+P////X/AADt/+j/AAAJAAAA+f8IABkAHgAbAA8ABQAUABAA7P/n//z/8v/l//L/+v/1//H/9f/7/wQA+P/i//f/DQAIAAAA/f8FAAkABQACAAUABQD6//3/CwAYAAQA8f///w0ACQD3//T/9P8BAAMA6P/w//r//f/x//b/BwD7//3/CQAYABAABQAEAAkAEgD8//j/+P/s//b/+//+//b/8f8HABkAFQABAPz/EgAdAA0A7//r/w0AEgD+//H/7f8BAA8ADwD8//T/CwAiABkA///8/wAABQD+//L/7v/q/+r/+v8NAA0AAgAGABgAFgAKAAYABAD1//D/9v/9//T/6P/1/wEAEQAZAAMA9/8IAA4ACgADAPb/6//w//z/+//5//b/9//7/wgAGgANAPn/+P8HABQADgD9/+7/CwAZAAUABAD1/+b/+f8DAPL/5P/t//H/9P8HAAkAAQATACEAHQAdAB4AGAALAP3/9//s/+H/2//e/+P/7f/2//z///8TACUAHQATABEAEwAMAAsACgAAAPX/+//8/wAACQD1/+T/5v/x//f/8f/Z/9j/+f8UABEACgAIAAwAHAAiABYA+P/y//n/AAD9/+D/5v8AAAYACwAHAAIACQAOAAkA/v/2//f/BgAGAPf/9/8BAAEA/v8BAP3/+P/y//D/AAASAAwAAwAFAB0AOAAhAAkADQALAAEA6//W/9D/0f/O/9j//f8DAAEAHAA2ADgAKwAdAA0ADAD5/+P/1P/S/97/4//r/+j/8P/4/wIACwAJAAEA/f8CAAcACwAMAAQA//8EAA0ADAD+//D/6v/2////+P/y//b//v8DAA4ADgAFAAAABwAQABAACAABAAAACgAUABAAFgAcABwAHQAYABgAGwASAAMA9v/0//D/6//s/+P/3v/r/wYAHAAlAC0AOgA9ADIAJwAbAP7/2f/E/7v/uP+y/7T/t//O/+r///8QAB8AKQAkACwAMwAsABcAAwAAAAMAAADw/9n/yP/Q/+L/4P/U/9j/7P8GABwAHgAoADMAMgA1ADYAKgAPAAIABgALAPz/4f/d//T/AAD1/+7/+v8KAAYACgAVABIAFQAfACEAGwATAAoABwD9//X/8v/l/9X/1v/m/+j/3//g/+n/+v8GAAIA9//u//v/CgADAPH/6v/v//H/7P/e/8//yP/U/9X/1v/o/wAADwAYACYAOQBIAEAAKgAZABsAFQADAO3/6P/9/w4AFQATACUAOwBBADcALgAhABYADgAAAPL/5v/p//H/8P/y/wIAFAASAAcADQAbABgACAD4//X/+//3/+n/3f/Q/8n/yf/D/73/wv/H/8z/1f/Y/+H/7//4//r/AAD0/+v/+P/8//X/5//f/+b/+P8EAAUAAwAAABYAPABGADEAJwA3AEMAQAAsABsAGgAcABwAFwAVAA8ADwAVACMALgAjAB0AIwAnACkAIAANAAgABQACAAcAAgD0/+z/7//x/+f/1//H/77/vP+3/7b/tv+t/6//v//T/9b/1P/j//P/+//6/+7/5f/p/+r/4P/S/87/2f/j/+b/6f/2/wcAFAAdACgAMwBCAE4A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a/98/4X/rP/p/xMAJQBJAG0AewCIAI4AfgBfADsAHwAGAN//tP+S/4j/hv+C/4//kv+l/8v/4v8BABIAIgA4AEUAVQBMAEYARAA1ADoANAAnACMAHAApACwAKQAxAC4AMwA3ADoAOAAnACUAIgAVABQADwAHAAYADwAaACEAJwAvADkAQgBKAEsARwBAADYAMAAnABgACwD+//v/+v/2//f/9P/4//z/AAAEAPr/9v/x/+b/3//P/7j/qv+e/5L/iP96/23/YP9g/2b/af+C/47/m//B/+f/DgAlADkAXgBrAHQAegBsAFgAMQAfAA4A4f/D/6j/m/+Z/5D/mv+k/6//x//h//r/CwAXACgANQBBAEEAOAA9ADcALwA0ACoAJgAoACUAKAAtADQALwArADUANAAvACsAHQAcABYAEAAOAAgACwALABUAIgAoADAANQBBAEoATABJAEAAOQAzACgAGQAHAAAA+P/y//L/8f/1//f/AAAJAAwADwAMAAwABwD+/+z/2v/M/7j/o/+S/4L/cv9o/1//VP9N/13/cf+A/5f/rP/a/wMAIgBBAFoAbwB3AIIAggBnAEcAKwATAPv/1f+t/53/mP+S/4//mv+m/7L/zv/t/wEACwAbADAAOwBCAEEAOQA2ADQAMQAlAB4AHgAZABwAIgAkACcAKgAuADUANQA1AC0AJgAlAB4AGQAPAAoACgAHAA4AFQAYAB0AJgAwADgAPAA/ADsAOQA4AC4AJgAZABEACwAGAAMAAQAAAAEACAAKAAsACwALAAYAAQD4/+z/2v/L/7v/qP+c/4v/f/93/3X/Z/9k/2j/df+G/5P/r/+9/9//CgAgADYASABiAGwAZwByAGAAQgAtACAADADi/83/v/+v/6//qf+n/7D/vP/N/9n/7v/9/wIAFQAkACsAKgAvADEALgAuAC4AKAAjACYAKAAlACkAKwAoACoALwAwACYAJAAjABoAFQAUAAsABwAJAAoADgAPABcAGwAgACoALgAvADAAMAAwACwAJwAhABcAFwAUABIADgAKABAAEgAUABoAFQAYABgAFQAPAAEA/P/r/+D/zv+9/6r/nP+T/4T/ev9y/2v/Yf9m/2n/eP+I/5n/rP/J//j/CQAmAEYAXABrAHMAgQB2AF4ATwA7ACEAAADj/8n/t/+q/5//m/+c/6f/sf/F/9j/5f/6/woAHAAmACoAMAAwAC8ALQAkACAAHAAXABQAEwAbABoAGgAiACgAMAAuAC8AMgAuAC4AKQAkAB4AGAAXABQAEQARAA8AEQAVABcAGQAaAB4AHwAjACQAIgAkACIAJAAlACMAIwAhACMAIwAiACIAHgAbABcAEQAJAP7/8//l/9n/y/++/7H/pf+b/5b/jv+L/4f/gv+E/33/g/+T/5z/pf+0/9D/6P/+/xcALQBBAE4AXwBnAGIAVwBFADoAJwAMAPL/3P/O/7v/sv+v/6r/r/+1/8b/1f/f//P//P8OABkAHgAkACEAKQAhABwAHAAVABEADwAPABQAFwAbAB8AJQA1ADYANQA5ADsAPQAzAC4AKwAgABsAFQASAA4ABgAIAAoADgAPAA8AGAAdACEAJQAnACoAKgAtAC0AKgAoACQAJAAhABwAFwATAA8ACQACAP//9f/s/+X/3f/V/8f/vv+1/63/pP+c/5b/iv+M/4f/ff+E/4r/nP+i/6z/yv/g//r/DgAkADkARQBUAGAAYABTAEQAPAAuABUA///l/9r/zf+//7r/s/+4/7n/x//U/9z/6f/0/wQADgAWABkAGgAfAB4AGwAYABUAEQARABEAEwARABUAHAAfACcAKwAxADEAMwA7ADcAMwAvACwAKgAjAB0AGQATABMAEAAPABEAEAATABUAHAAfAB4AIgAlACcAJwAjACYAIwAeABwAFwAUAA8ACwAHAAEA///5//P/7v/n/+D/2f/S/8r/v/+5/7P/qv+k/5z/mP+W/5D/kP+X/5//rf+0/8L/3f/t////EQAgADAAMwA/AEQAOAAxACcAIgAVAAAA9//q/+H/3f/R/9H/0f/U/9z/3f/o//D/9/8BAAcADQASABUAGQAbABwAGAAXABkAGAAXABUAFwAaABwAHgAhACUAJwArACwALQAvAC0ALQAsACoAKQAkACMAIQAdABwAGAAXABcAFAATABEAFAATABAAEgASABQAEwAUABUAEgATABIAEQAOAAsACAAEAAAA+//0/+7/6P/h/9n/0//O/8X/vf+7/7X/rv+q/6b/o/+e/6T/qf+t/7X/vv/S/93/7////wcAGwAkAC8ANAAyADEALAAkABoACgD9//P/5v/f/9X/0f/P/9H/2v/d/+b/7v/8/wcADAAYABoAHwAiAB8AHgAZABUAEgAPAAsACgALAA8AEQAXAB0AIwApAC0ANQA2ADcANwA0ADMALAAmACAAHAAXAA8ADAALAAsACQAJAA4ADgARABQAFwAYABYAGAAWABQAEgAMAAkABQAEAAEA/v/8//v/+v/6//j/9v/z//H/7v/o/+L/2v/S/8z/xP+8/7L/rP+o/6T/ov+n/67/s/+8/8z/2//p//j/BQATABsAIgApACcAJQAeABcAEAADAPr/8f/n/+P/3//b/93/3v/m/+r/8P/7/wIACwAPABUAGgAbABsAGQAXABQADwAOAAsABwAHAAgADAAOABEAGAAcACIAKAAsAC8ALwAxADEALQArACgAJAAgABkAFwAUABEADwAMAA4ADQAMAA0ADQAPAA8ADwAPAAwACwALAAgABgAEAAEAAAD///7//P/7//r/+P/3//X/8f/t/+j/5f/f/9j/0v/L/8f/wv+9/7r/uf+7/7//w//J/9D/2f/i/+z/9v/8/wAACAAOAA8ADgANAAwABwAEAAAA+v/1//L/8f/v/+7/8P/z//f//f8AAAYADAARABYAGAAbABwAGwAaABcAFAARAA4ADQAKAAkACgALAAwADwATABcAGwAfACMAJgAoACoAKwArACoAJwAkACIAHQAYABUAEQAMAAgABgADAAEAAAD//////v/+//3//v/+//7//f/9//7//P/7//v/+f/4//b/9f/z//H/7v/s/+v/6P/m/+T/4v/i/+H/4P/g/97/3//g/+H/4f/i/+P/5f/o/+r/7P/u//H/9f/3//r//P/9////AAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAgADAAQABgAIAAoADAAPAA8AEAASABIAEgARABAADwAPAA0ADQALAAoACgAKAAoACgAKAAoACwAMAAsADAAMAA0ADgANAAwADAAMAAsACgAJAAgABwAHAAYABgAGAAUABQAFAAUABQAEAAQABAACAAEAAAD///3/+//6//j/9//1//T/8//y//H/8v/x//L/8v/z//T/9f/1//b/9//3//j/9//4//j/9//3//b/9f/0//T/8v/y//H/8v/y//L/8v/y//T/9f/2//f/+P/6//z//v///wAAAQADAAQABQAFAAYABgAHAAcABQAFAAUABQAFAAUABQAGAAcABwAHAAgACQAJAAkACQAJAAkACQAJAAkACQAJAAkACQAKAAoACgAKAAoACgAJAAkACAAHAAYABQAEAAMAAgABAAAAAAAAAP/////+//7//v/+//7////////////////////+//7//f/8//z/+//6//r/+v/6//r/+//7//3//f/+//7//v/+//7//v/9//3//P/7//v/+v/6//r/+v/6//r/+//8//z//f/9//7///8AAAAAAAAAAAEAAgACAAMAAwAEAAQABAAEAAMABAADAAMAAgACAAEAAQABAAEAAAAAAAAAAAAAAAEAAQABAAEAAQABAAEAAQACAAEAAQACAAIAAgACAAIAAgACAAIAAgACAAIAAgACAAEAAQABAAAAAAAAAAAA///////////+//7//v////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP////////7//v/+//7//v/+//////8AAAAAAAAAAAEAAQABAAEAAQABAAEAAQABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAEAAAABAAEAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP///////wAA/////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAQABAAEAAQABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////////////////////////////////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//AAAAAP//////////////////AAD//wAAAAAAAAAAAAAAAAAAAAAAAP///////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP//////////AAAAAP//AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAQABAAEAAQABAAEAAQABAAEAAQABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//8AAAAAAAAAAAAAAAAAAAAAAQAAAAAAAQAAAAEAAQAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//8AAP///////wAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAD/////////////AAAAAP//AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAEAAAABAAEAAQABAAAAAQABAAAAAAAAAAAAAAAAAP//AAAAAP///////wAAAAAAAAAAAAAAAAEAAQAAAAAAAgADAAIAAQACAAMAAgABAAAAAAAAAAAAAAD///////8AAAAAAAAAAAAAAAAAAAAAAAAAAP////////7/AAAAAAAA/////wAAAAAAAAAAAAAAAAAAAAAAAAAA///////////+//7//v8AAAAAAAD//wAAAAABAAIAAAAAAAIAAgACAAAAAAABAAEAAAAAAAAA//8AAP//AAAAAAAAAAAAAAAAAQABAAAAAAAAAAAAAAAAAAAA"

	// 解码Base64字符串
	//pcmData, err := base64.StdEncoding.DecodeString(base64String)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 假设你想将PCM数据转换为WAV格式
	wavData := base64String

	// 创建WAV文件
	wavFile, err := os.Create("./output1.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer wavFile.Close()

	// 将PCM数据写入WAV文件
	if _, err := wavFile.Write([]byte(wavData)); err != nil {
		log.Fatal(err)
	}

	fmt.Println("PCM数据已成功转换为WAV文件")
}