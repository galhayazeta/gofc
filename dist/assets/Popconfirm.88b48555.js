import{r as z,w as N,c as B,x as S,a9 as _,A as R,$ as m,u as T,p as L,z as g,y as h,q as I,N as $,s as M,e as P,j as u,d as C,m as j,K as q,ab as E,aq as F,t as x}from"./Button.f622f226.js";import{h as a,i as w,k as K,j as f,t as b,d as U,p as V}from"./index.08bac0c2.js";function W(e,o=[],n){const t={};return Object.getOwnPropertyNames(e).forEach(r=>{o.includes(r)||(t[r]=e[r])}),Object.assign(t,n)}var Z=z("warning",a("svg",{viewBox:"0 0 24 24",version:"1.1",xmlns:"http://www.w3.org/2000/svg"},a("g",{stroke:"none","stroke-width":"1","fill-rule":"evenodd"},a("g",{"fill-rule":"nonzero"},a("path",{d:"M12,2 C17.523,2 22,6.478 22,12 C22,17.522 17.523,22 12,22 C6.477,22 2,17.522 2,12 C2,6.478 6.477,2 12,2 Z M12.0018002,15.0037242 C11.450254,15.0037242 11.0031376,15.4508407 11.0031376,16.0023869 C11.0031376,16.553933 11.450254,17.0010495 12.0018002,17.0010495 C12.5533463,17.0010495 13.0004628,16.553933 13.0004628,16.0023869 C13.0004628,15.4508407 12.5533463,15.0037242 12.0018002,15.0037242 Z M11.99964,7 C11.4868042,7.00018474 11.0642719,7.38637706 11.0066858,7.8837365 L11,8.00036004 L11.0018003,13.0012393 L11.00857,13.117858 C11.0665141,13.6151758 11.4893244,14.0010638 12.0021602,14.0008793 C12.514996,14.0006946 12.9375283,13.6145023 12.9951144,13.1171428 L13.0018002,13.0005193 L13,7.99964009 L12.9932303,7.8830214 C12.9352861,7.38570354 12.5124758,6.99981552 11.99964,7 Z"}))))),A={iconSize:"22px"};const D=e=>{const{fontSize:o,warningColor:n}=e;return Object.assign(Object.assign({},A),{fontSize:o,iconColor:n})},H=N({name:"Popconfirm",common:B,peers:{Button:S,Popover:_},self:D});var G=H;const y=R("n-popconfirm"),O={positiveText:String,negativeText:String,showIcon:{type:Boolean,default:!0},onPositiveClick:{type:Function,required:!0},onNegativeClick:{type:Function,required:!0}},k=M(O);var J=w({name:"NPopconfirmPanel",props:O,setup(e){const{localeRef:o}=m("Popconfirm"),{inlineThemeDisabled:n}=T(),{mergedClsPrefixRef:t,mergedThemeRef:c,props:r}=K(y),d=f(()=>{const{common:{cubicBezierEaseInOut:s},self:{fontSize:p,iconSize:v,iconColor:l}}=c.value;return{"--n-bezier":s,"--n-font-size":p,"--n-icon-size":v,"--n-icon-color":l}}),i=n?L("popconfirm-panel",void 0,d,r):void 0;return Object.assign(Object.assign({},m("Popconfirm")),{mergedClsPrefix:t,cssVars:n?void 0:d,localizedPositiveText:f(()=>e.positiveText||o.value.positiveText),localizedNegativeText:f(()=>e.negativeText||o.value.negativeText),positiveButtonProps:b(r,"positiveButtonProps"),negativeButtonProps:b(r,"negativeButtonProps"),handlePositiveClick(s){e.onPositiveClick(s)},handleNegativeClick(s){e.onNegativeClick(s)},themeClass:i==null?void 0:i.themeClass,onRender:i==null?void 0:i.onRender})},render(){var e;const{mergedClsPrefix:o,showIcon:n,$slots:t}=this,c=g(t.action,()=>this.negativeText===null&&this.positiveText===null?[]:[this.negativeText!==null&&a(h,Object.assign({size:"small",onClick:this.handleNegativeClick},this.negativeButtonProps),{default:()=>this.localizedNegativeText}),this.positiveText!==null&&a(h,Object.assign({size:"small",type:"primary",onClick:this.handlePositiveClick},this.positiveButtonProps),{default:()=>this.localizedPositiveText})]);return(e=this.onRender)===null||e===void 0||e.call(this),a("div",{class:[`${o}-popconfirm__panel`,this.themeClass],style:this.cssVars},I(t.default,r=>n||r?a("div",{class:`${o}-popconfirm__body`},n?a("div",{class:`${o}-popconfirm__icon`},g(t.icon,()=>[a($,{clsPrefix:o},{default:()=>a(Z,null)})])):null,r):null),c?a("div",{class:[`${o}-popconfirm__action`]},c):null)}}),Q=P("popconfirm",[u("body",`
 font-size: var(--n-font-size);
 display: flex;
 align-items: center;
 flex-wrap: nowrap;
 position: relative;
 `,[u("icon",`
 display: flex;
 font-size: var(--n-icon-size);
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 margin: 0 8px 0 0;
 `)]),u("action",`
 display: flex;
 justify-content: flex-end;
 `,[C("&:not(:first-child)","margin-top: 8px"),P("button",[C("&:not(:last-child)","margin-right: 8px;")])])]);const X=Object.assign(Object.assign(Object.assign({},j.props),F),{positiveText:String,negativeText:String,showIcon:{type:Boolean,default:!0},trigger:{type:String,default:"click"},positiveButtonProps:Object,negativeButtonProps:Object,onPositiveClick:Function,onNegativeClick:Function});var oe=w({name:"Popconfirm",props:X,__popover__:!0,setup(e){const{mergedClsPrefixRef:o}=T(),n=j("Popconfirm","-popconfirm",Q,G,e,o),t=U(null);function c(i){const{onPositiveClick:s,"onUpdate:show":p}=e;Promise.resolve(s?s(i):!0).then(v=>{var l;v!==!1&&((l=t.value)===null||l===void 0||l.setShow(!1),p&&x(p,!1))})}function r(i){const{onNegativeClick:s,"onUpdate:show":p}=e;Promise.resolve(s?s(i):!0).then(v=>{var l;v!==!1&&((l=t.value)===null||l===void 0||l.setShow(!1),p&&x(p,!1))})}return V(y,{mergedThemeRef:n,mergedClsPrefixRef:o,props:e}),Object.assign(Object.assign({},{setShow(i){var s;(s=t.value)===null||s===void 0||s.setShow(i)},syncPosition(){var i;(i=t.value)===null||i===void 0||i.syncPosition()}}),{mergedTheme:n,popoverInstRef:t,handlePositiveClick:c,handleNegativeClick:r})},render(){const{$slots:e,$props:o,mergedTheme:n}=this;return a(E,W(o,k,{theme:n.peers.Popover,themeOverrides:n.peerOverrides.Popover,internalExtraClass:["popconfirm"],ref:"popoverInstRef"}),{trigger:e.activator||e.trigger,default:()=>{const t=q(o,k);return a(J,Object.assign(Object.assign({},t),{onPositiveClick:this.handlePositiveClick,onNegativeClick:this.handleNegativeClick}),e)}})}});export{oe as N,Z as W};
