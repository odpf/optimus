"use strict";(self.webpackChunkoptimus=self.webpackChunkoptimus||[]).push([[3807],{3905:function(e,t,n){n.d(t,{Zo:function(){return p},kt:function(){return d}});var i=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,i)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,i,a=function(e,t){if(null==e)return{};var n,i,a={},r=Object.keys(e);for(i=0;i<r.length;i++)n=r[i],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(i=0;i<r.length;i++)n=r[i],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=i.createContext({}),u=function(e){var t=i.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},p=function(e){var t=u(e.components);return i.createElement(s.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},c=i.forwardRef((function(e,t){var n=e.components,a=e.mdxType,r=e.originalType,s=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),c=u(n),d=a,g=c["".concat(s,".").concat(d)]||c[d]||m[d]||r;return n?i.createElement(g,l(l({ref:t},p),{},{components:n})):i.createElement(g,l({ref:t},p))}));function d(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var r=n.length,l=new Array(r);l[0]=c;var o={};for(var s in t)hasOwnProperty.call(t,s)&&(o[s]=t[s]);o.originalType=e,o.mdxType="string"==typeof e?e:a,l[1]=o;for(var u=2;u<r;u++)l[u]=n[u];return i.createElement.apply(null,l)}return i.createElement.apply(null,n)}c.displayName="MDXCreateElement"},673:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return o},contentTitle:function(){return s},metadata:function(){return u},toc:function(){return p},default:function(){return c}});var i=n(7462),a=n(3366),r=(n(7294),n(3905)),l=["components"],o={},s=void 0,u={unversionedId:"rfcs/simplify_plugin_maintenance",id:"rfcs/simplify_plugin_maintenance",isDocsHomePage:!1,title:"simplify_plugin_maintenance",description:"- Feature Name: Simplify Plugins",source:"@site/docs/rfcs/20220507_simplify_plugin_maintenance.md",sourceDirName:"rfcs",slug:"/rfcs/simplify_plugin_maintenance",permalink:"/optimus/docs/rfcs/simplify_plugin_maintenance",editUrl:"https://github.com/odpf/optimus/edit/master/docs/docs/rfcs/20220507_simplify_plugin_maintenance.md",tags:[],version:"current",lastUpdatedBy:"Arinda Arif",lastUpdatedAt:1657265432,formattedLastUpdatedAt:"7/8/2022",sidebarPosition:20220507,frontMatter:{}},p=[{value:"Background :",id:"background-",children:[{value:"Changes that trigger a new release in Optimus setup:",id:"changes-that-trigger-a-new-release-in-optimus-setup",children:[]},{value:"Release dependencies as per current design",id:"release-dependencies-as-per-current-design",children:[]},{value:"1. <u>Avoid Wrapping Executor Images</u>  :",id:"1-avoid-wrapping-executor-images--",children:[]},{value:"2. <u>Simplify Plugin Installation</u> :",id:"2-simplify-plugin-installation-",children:[]}]},{value:"Approach :",id:"approach-",children:[{value:"1. <u>Avoid Wrapping Executor Images </u> :",id:"1-avoid-wrapping-executor-images---1",children:[]},{value:"2. <u>Simplify Plugin Installations</u> :",id:"2-simplify-plugin-installations-",children:[]}]},{value:"Result:",id:"result",children:[]}],m={toc:p};function c(e){var t=e.components,n=(0,a.Z)(e,l);return(0,r.kt)("wrapper",(0,i.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Feature Name: Simplify Plugins"),(0,r.kt)("li",{parentName:"ul"},"Status: Draft"),(0,r.kt)("li",{parentName:"ul"},"Start Date: 2022-05-07"),(0,r.kt)("li",{parentName:"ul"},"Author: Saikumar")),(0,r.kt)("h1",{id:"summary"},"Summary"),(0,r.kt)("p",null,"The scope of this rfc is to simplify the release and deployment operations w.r.t the optimus plugin ecosystem."),(0,r.kt)("p",null,"The proposal here is to :"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("strong",{parentName:"li"},"Avoid Wrapping Executor Images")," :",(0,r.kt)("br",{parentName:"li"}),"Decouple the executor_boot_process and the executor as separate containers where the airflow worker launches a pod with init-container (for boot process) adjacent to executor container."),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("strong",{parentName:"li"},"Simplfy Plugin Installation")," :",(0,r.kt)("ul",{parentName:"li"},(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"Server end")," : Install plugins on-demand declaratively instead of manually baking them into the optimus server image (in kubernetes setup)."),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"Client end")," : Plugin interface for client-end is limited to support  Version, Survey Questions and Answers etc. that can be extracted out from the current plugin interface and maintained as yaml file which simplifies platform dependent plugin distribution for cli.")))),(0,r.kt)("h1",{id:"technical-design"},"Technical Design"),(0,r.kt)("h2",{id:"background-"},"Background :"),(0,r.kt)("h3",{id:"changes-that-trigger-a-new-release-in-optimus-setup"},"Changes that trigger a new release in Optimus setup:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Executor Image changes"),(0,r.kt)("li",{parentName:"ul"},"Executor Image Wrapper changes"),(0,r.kt)("li",{parentName:"ul"},"Plugin binary changes"),(0,r.kt)("li",{parentName:"ul"},"Optimus binary changes")),(0,r.kt)("h3",{id:"release-dependencies-as-per-current-design"},"Release dependencies as per current design"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"Executor Image release")," -> ",(0,r.kt)("inlineCode",{parentName:"li"},"Executor Wrapper Image release")," -> ",(0,r.kt)("inlineCode",{parentName:"li"},"Plugin binary release")," -> ",(0,r.kt)("inlineCode",{parentName:"li"},"Server release")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"Plugin binary release")," -> ",(0,r.kt)("inlineCode",{parentName:"li"},"Server release"))),(0,r.kt)("h3",{id:"1-avoid-wrapping-executor-images--"},"1. ",(0,r.kt)("u",null,"Avoid Wrapping Executor Images"),"  :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"The ",(0,r.kt)("inlineCode",{parentName:"li"},"executor_boot_process")," and ",(0,r.kt)("inlineCode",{parentName:"li"},"executor")," are coupled:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"-- Plugin repo structure\n/task/\n/task/Dockerfile           -- task_image\n/task/executor/Dockerfile  -- executor_image\n")),(0,r.kt)("p",null,"Executor Wrapper Image (Task Image) :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"It's a wrapper around the executor_image to facilitate boot mechanism for executor."),(0,r.kt)("li",{parentName:"ul"},"The optimus binary is downloaded during buildtime of this image."),(0,r.kt)("li",{parentName:"ul"},"During runtime, it does as follow :",(0,r.kt)("ul",{parentName:"li"},(0,r.kt)("li",{parentName:"ul"},"Fetch assets, secrets, env from optimus server."),(0,r.kt)("li",{parentName:"ul"},"Load the env and launches the executor process.")))),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"task_image \n    | executor_image\n    | optimus-bin\n    | entrypoint.sh (load assets, env and launch executor)\n")),(0,r.kt)("p",null,"The ",(0,r.kt)("inlineCode",{parentName:"p"},"optimus-bin")," and ",(0,r.kt)("inlineCode",{parentName:"p"},"entrypoint.sh")," are baked into the ",(0,r.kt)("inlineCode",{parentName:"p"},"task_image")," and is being maintained by task/plugin developers."),(0,r.kt)("h3",{id:"2-simplify-plugin-installation-"},"2. ",(0,r.kt)("u",null,"Simplify Plugin Installation")," :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Plugin binaries are manually installed (baked into optimus image in kubernetes setup). "),(0,r.kt)("li",{parentName:"ul"},"Any change in plugin code demands re-creation of optimus image with new plugin binary, inturn demanding redeployment of optimus server. (in kubernetes setup)"),(0,r.kt)("li",{parentName:"ul"},"At client side, plugin binaries require support for different platforms.")),(0,r.kt)("h2",{id:"approach-"},"Approach :"),(0,r.kt)("h3",{id:"1-avoid-wrapping-executor-images---1"},"1. ",(0,r.kt)("u",null,"Avoid Wrapping Executor Images ")," :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Decouple the lifecycle of the executor and the boot process as seperate containers/images.")),(0,r.kt)("img",{src:"images/simplify_plugins_executor.png",alt:"Simplify Plugins Executor",width:"800"}),(0,r.kt)("p",null,(0,r.kt)("strong",{parentName:"p"},"Task Boot Sequence"),":"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Airflow worker fetches env and secrets for the job and adds them to the executor pod as environment variables."),(0,r.kt)("li",{parentName:"ol"},"KubernetesPodOperator spawns init-container and executor-container, mounted with shared volume (type emptyDir) for assets."),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("inlineCode",{parentName:"li"},"init-container")," fetches assets, config, env files and writes onto the shared volume."),(0,r.kt)("li",{parentName:"ol"},"the default entrypoint in the executor-image starts the actual job.")),(0,r.kt)("h3",{id:"2-simplify-plugin-installations-"},"2. ",(0,r.kt)("u",null,"Simplify Plugin Installations")," :"),(0,r.kt)("img",{src:"images/plugin_manager.png",alt:"Plugins Manager",width:"800"}),(0,r.kt)("h4",{id:"a-plugin-manager-at-server-end"},"A) Plugin Manager: (at server end)"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Currently the plugins are maintained as mono repo and internally versioned i.e., version of plugin are mainatined within each plugin."),(0,r.kt)("li",{parentName:"ul"},"A plugin manager is required to support declarative installation of plugins."),(0,r.kt)("li",{parentName:"ul"},"This plugin manager consumes a config (plugin_manager_config) and downloads artifacts from a plugin repository.")),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Optimus support for plugin manager as below.",(0,r.kt)("ul",{parentName:"li"},(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"optimus plugin install -c optimus.yaml")," -- server"),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"optimus plugin sync -c optimus.yaml")," -- cli"),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"optimus plugin list")))),(0,r.kt)("li",{parentName:"ul"},"Support for different kinds of plugin repositories (like s3, gcs, url, local file system etc..) gives the added flexibility and options to distribute and install the plugin binaries in different ways."),(0,r.kt)("li",{parentName:"ul"},"Plugins are installed at container runtime and this decouples the building of optimus docker image from plugins installations. The plugin_manager_config can be maintained as ",(0,r.kt)("inlineCode",{parentName:"li"},"ConfigMap")," so as to reflect any updates in plugins all one needs to do is change in the config map and restart the pod. "),(0,r.kt)("li",{parentName:"ul"},"Example for the plugin config: ")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'    plugins :\n      plugin_dir : ""\n      providers :\n      - type : http\n        name : internal_url_xyz_org\n        url: http://<internal_url>\n        auth: \n      - type : gcs\n        name : private_gcs_backend_team\n        bucket: <bucket>\n        service_account : <base64_encoded_service_account>\n      plugins :\n      - provider : internal_url_xyz_org\n        path : <plugin_name>.tar.gz\n      - provider : private_gcs_backend_team\n        path : <plugin_name>.zip\n        .\n        .\n      \n')),(0,r.kt)("h4",{id:"b-plugin-yaml-interface-for-client-side-simplification"},"B) Plugin Yaml Interface: (for client side simplification)"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Currently plugins are implemented and distributed as binaries and as clients needs to install them, it demands support for different host  architectures."),(0,r.kt)("li",{parentName:"ul"},"Since CLI (client side) plugins just require details about plugin such as Version, Suevery Questions etc. the proposal here is to maintain CLI plugins as yaml files."),(0,r.kt)("li",{parentName:"ul"},"Implementation wise, the proposal here is to split the current plugin interface (which only supports interaction with binary plugins) to also accommodate yaml based plugins."),(0,r.kt)("li",{parentName:"ul"},"The above mentioned pluign manager, at server end, would be agnostic about the contents of plugin artifacts from the repository."),(0,r.kt)("li",{parentName:"ul"},"At client side, the CLI could sync the yaml files from the server itself to stay up-to-date with the server wrt plugins.")),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Example representation of the yaml plugin : ")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'  Name: bq2bq\n  Version: latest\n  Image: docker.io/odpf/optimus-task-bq2bq\n  Description: "BigQuery to BigQuery transformation task"\n  Questions:\n    - Name:    "Project"\n      Prompt:  "Project ID"\n      Help:    "Destination bigquery project ID"\n    - Name:\n      .\n      .\n      .\n  DefaultConfig:\n    - name: PROJECT\n      value: \'\'\n    - name: TABLE\n      value: \'\'\n\n  DefaultAssets:\n    - name: query.sql\n      value: |\n        -- SQL query goes here\n        -- Select * from "project.dataset.table"\n      \n')),(0,r.kt)("h2",{id:"result"},"Result:"),(0,r.kt)("img",{src:"images/simplify_plugins.png",alt:"Simplify Plugins",width:"800"}),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Executor boot process is standardised and extracted away from plugin developers. Now any arbitrary image can be used for executors."),(0,r.kt)("li",{parentName:"ul"},"At server side, for changes in plugin (dur to plugin release), update the plugin_manager_config and restart the optimus server pod. The plugin manager is expected to reinstall the plugins."),(0,r.kt)("li",{parentName:"ul"},"Client side dependency on plugins is simplified with yaml based plugins.")))}c.isMDXComponent=!0}}]);