"use strict";(self.webpackChunkoptimus=self.webpackChunkoptimus||[]).push([[6886],{3905:function(e,t,n){n.d(t,{Zo:function(){return l},kt:function(){return m}});var i=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,i)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,i,r=function(e,t){if(null==e)return{};var n,i,r={},s=Object.keys(e);for(i=0;i<s.length;i++)n=s[i],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(i=0;i<s.length;i++)n=s[i],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var c=i.createContext({}),u=function(e){var t=i.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},l=function(e){var t=u(e.components);return i.createElement(c.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},d=i.forwardRef((function(e,t){var n=e.components,r=e.mdxType,s=e.originalType,c=e.parentName,l=o(e,["components","mdxType","originalType","parentName"]),d=u(n),m=r,h=d["".concat(c,".").concat(m)]||d[m]||p[m]||s;return n?i.createElement(h,a(a({ref:t},l),{},{components:n})):i.createElement(h,a({ref:t},l))}));function m(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var s=n.length,a=new Array(s);a[0]=d;var o={};for(var c in t)hasOwnProperty.call(t,c)&&(o[c]=t[c]);o.originalType=e,o.mdxType="string"==typeof e?e:r,a[1]=o;for(var u=2;u<s;u++)a[u]=n[u];return i.createElement.apply(null,a)}return i.createElement.apply(null,n)}d.displayName="MDXCreateElement"},4730:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return o},contentTitle:function(){return c},metadata:function(){return u},toc:function(){return l},default:function(){return d}});var i=n(7462),r=n(3366),s=(n(7294),n(3905)),a=["components"],o={},c="Architecture",u={unversionedId:"concepts/architecture",id:"concepts/architecture",isDocsHomePage:!1,title:"Architecture",description:"Basic building blocks of Optimus are",source:"@site/docs/concepts/architecture.md",sourceDirName:"concepts",slug:"/concepts/architecture",permalink:"/optimus/docs/concepts/architecture",editUrl:"https://github.com/odpf/optimus/edit/master/docs/docs/concepts/architecture.md",tags:[],version:"current",lastUpdatedBy:"Yash Bhardwaj",lastUpdatedAt:1658204053,formattedLastUpdatedAt:"7/19/2022",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Overview",permalink:"/optimus/docs/concepts/overview"},next:{title:"Understanding Intervals and Windows",permalink:"/optimus/docs/concepts/intervals-and-windows"}},l=[{value:"Overview",id:"overview",children:[]},{value:"Optimus CLI",id:"optimus-cli",children:[]},{value:"Optimus Service",id:"optimus-service",children:[]},{value:"Optimus Database",id:"optimus-database",children:[]},{value:"Optimus Plugins",id:"optimus-plugins",children:[]},{value:"Scheduler",id:"scheduler",children:[]},{value:"Optimus Extension",id:"optimus-extension",children:[]}],p={toc:l};function d(e){var t=e.components,o=(0,r.Z)(e,a);return(0,s.kt)("wrapper",(0,i.Z)({},p,o,{components:t,mdxType:"MDXLayout"}),(0,s.kt)("h1",{id:"architecture"},"Architecture"),(0,s.kt)("p",null,"Basic building blocks of Optimus are"),(0,s.kt)("ul",null,(0,s.kt)("li",{parentName:"ul"},"Optimus CLI"),(0,s.kt)("li",{parentName:"ul"},"Optimus Service"),(0,s.kt)("li",{parentName:"ul"},"Optimus Database"),(0,s.kt)("li",{parentName:"ul"},"Optimus Plugins"),(0,s.kt)("li",{parentName:"ul"},"Scheduler"),(0,s.kt)("li",{parentName:"ul"},"Optimus Extension")),(0,s.kt)("h3",{id:"overview"},"Overview"),(0,s.kt)("p",null,(0,s.kt)("img",{alt:"Architecture Diagram",src:n(4312).Z,title:"OptimusArchitecture"})),(0,s.kt)("h3",{id:"optimus-cli"},"Optimus CLI"),(0,s.kt)("p",null,(0,s.kt)("inlineCode",{parentName:"p"},"optimus")," is a command line tool used to interact with the main optimus service and basic scaffolding job\nspecifications. It can be used to"),(0,s.kt)("ul",null,(0,s.kt)("li",{parentName:"ul"},"Generate jobs based on user inputs"),(0,s.kt)("li",{parentName:"ul"},"Add hooks to existing jobs"),(0,s.kt)("li",{parentName:"ul"},"Dump a compiled specification for the consumption of a scheduler"),(0,s.kt)("li",{parentName:"ul"},"Deployment of specifications to ",(0,s.kt)("inlineCode",{parentName:"li"},"Optimus Service")),(0,s.kt)("li",{parentName:"ul"},"Create resource specifications for datastores"),(0,s.kt)("li",{parentName:"ul"},"Start optimus server")),(0,s.kt)("p",null,"Optimus also has an admin flag that can be turned on using ",(0,s.kt)("inlineCode",{parentName:"p"},"OPTIMUS_ADMIN_ENABLED=1")," env flag.\nThis hides few commands which are used internally during the lifecycle of tasks/hooks\nexecution."),(0,s.kt)("h3",{id:"optimus-service"},"Optimus Service"),(0,s.kt)("p",null,"Optimus cli can start a service that controls and orchestrates all that Optimus has to\noffer. Optimus cli uses GRPC to communicate with the optimus service for almost all the\noperations that takes ",(0,s.kt)("inlineCode",{parentName:"p"},"host")," as the flag. Service also exposes few REST endpoints\nthat can be used with simple curl request for registering a new project or checking\nthe status of a job, etc."),(0,s.kt)("p",null,"As soon as jobs are ready in a repository, a deployment request is sent to the service\nwith all the specs(normally in yaml) which are parsed and stored in the database.\nOnce these specs are stored, each of them are compiled to generate a scheduler parsable\njob format which will be eventually consumed by a supported scheduler to execute the\njob. These compiled specifications are uploaded to an ",(0,s.kt)("strong",{parentName:"p"},"object store")," which gets synced\nto the scheduler."),(0,s.kt)("h3",{id:"optimus-database"},"Optimus Database"),(0,s.kt)("p",null,"Specifications once requested for deployment needs to be stored somewhere as a source\nof truth. Optimus uses postgres as a storage engine to store raw specifications, job\nassets, run details, project configurations, etc."),(0,s.kt)("h3",{id:"optimus-plugins"},"Optimus Plugins"),(0,s.kt)("p",null,"Optimus itself doesn't govern how a job is supposed to execute the transformation. It\nonly provides the building blocks which needs to be implemented by a task. A plugin is\ndivided in two parts, an adapter and a docker image. Docker image contains the actual\ntransformation logic that needs to be executed in the task and adapter helps optimus\nto understand what all this task can do and help in doing it."),(0,s.kt)("h3",{id:"scheduler"},"Scheduler"),(0,s.kt)("p",null,"Job adapters consumes job specifications which eventually needs to be scheduled and\nexecuted via a execution engine. This execution engine is termed here as Scheduler.\nOptimus by default recommends using ",(0,s.kt)("inlineCode",{parentName:"p"},"Airflow")," but is extensible enough to support any\nother scheduler that satisfies some basic requirements, one of the most important\nof all is, scheduler should be able to execute a Docker container."),(0,s.kt)("h3",{id:"optimus-extension"},"Optimus Extension"),(0,s.kt)("p",null,"Optimus extension is a feature in Optimus where the user could extend the\nfunctionality of Optimus itself using third-party or arbitrary\nimplementation. Currently, extension is designed only for when the\nuser running it as CLI."))}d.isMDXComponent=!0},4312:function(e,t,n){t.Z=n.p+"assets/images/OptimusArchitecture_dark_07June2021-496f10b02b693c3113dd4800802b131e.png"}}]);