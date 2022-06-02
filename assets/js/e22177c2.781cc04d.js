"use strict";(self.webpackChunkoptimus=self.webpackChunkoptimus||[]).push([[3807],{3905:function(e,n,t){t.d(n,{Zo:function(){return p},kt:function(){return m}});var a=t(7294);function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function r(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);n&&(a=a.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,a)}return t}function l(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?r(Object(t),!0).forEach((function(n){i(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):r(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function o(e,n){if(null==e)return{};var t,a,i=function(e,n){if(null==e)return{};var t,a,i={},r=Object.keys(e);for(a=0;a<r.length;a++)t=r[a],n.indexOf(t)>=0||(i[t]=e[t]);return i}(e,n);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(a=0;a<r.length;a++)t=r[a],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var u=a.createContext({}),s=function(e){var n=a.useContext(u),t=n;return e&&(t="function"==typeof e?e(n):l(l({},n),e)),t},p=function(e){var n=s(e.components);return a.createElement(u.Provider,{value:n},e.children)},c={inlineCode:"code",wrapper:function(e){var n=e.children;return a.createElement(a.Fragment,{},n)}},d=a.forwardRef((function(e,n){var t=e.components,i=e.mdxType,r=e.originalType,u=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),d=s(t),m=i,g=d["".concat(u,".").concat(m)]||d[m]||c[m]||r;return t?a.createElement(g,l(l({ref:n},p),{},{components:t})):a.createElement(g,l({ref:n},p))}));function m(e,n){var t=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var r=t.length,l=new Array(r);l[0]=d;var o={};for(var u in n)hasOwnProperty.call(n,u)&&(o[u]=n[u]);o.originalType=e,o.mdxType="string"==typeof e?e:i,l[1]=o;for(var s=2;s<r;s++)l[s]=t[s];return a.createElement.apply(null,l)}return a.createElement.apply(null,t)}d.displayName="MDXCreateElement"},673:function(e,n,t){t.r(n),t.d(n,{frontMatter:function(){return o},contentTitle:function(){return u},metadata:function(){return s},toc:function(){return p},default:function(){return d}});var a=t(7462),i=t(3366),r=(t(7294),t(3905)),l=["components"],o={},u=void 0,s={unversionedId:"rfcs/simplify_plugin_maintenance",id:"rfcs/simplify_plugin_maintenance",isDocsHomePage:!1,title:"simplify_plugin_maintenance",description:"- Feature Name: Simplify Plugin Releases",source:"@site/docs/rfcs/20220507_simplify_plugin_maintenance.md",sourceDirName:"rfcs",slug:"/rfcs/simplify_plugin_maintenance",permalink:"/optimus/docs/rfcs/simplify_plugin_maintenance",editUrl:"https://github.com/odpf/optimus/edit/master/docs/docs/rfcs/20220507_simplify_plugin_maintenance.md",tags:[],version:"current",lastUpdatedBy:"Anwar Hidayat",lastUpdatedAt:1654166089,formattedLastUpdatedAt:"6/2/2022",sidebarPosition:20220507,frontMatter:{}},p=[{value:"Background :",id:"background-",children:[{value:"Optimus components that trigger a new release:",id:"optimus-components-that-trigger-a-new-release",children:[]},{value:"Non-trivial release dependencies as per current design",id:"non-trivial-release-dependencies-as-per-current-design",children:[]},{value:"1. <u>Dependency between <code>Executor and Task</code></u>  :",id:"1-dependency-between-executor-and-task--",children:[]},{value:"2. <u>Dependency between <code>Plugin and Server</code>, also  <code>Plugin and Executor</code></u> :",id:"2-dependency-between-plugin-and-server-also--plugin-and-executor-",children:[]}]},{value:"Approach :",id:"approach-",children:[{value:"1. <u>Removing dependency between <code>Executor and Task Image</code> releases</u> :",id:"1-removing-dependency-between-executor-and-task-image-releases-",children:[]},{value:"2. <u>Removing dependency between <code>Plugin and Server</code> releases</u> :",id:"2-removing-dependency-between-plugin-and-server-releases-",children:[]},{value:"3. <u>Removing dependency between <code>Plugin and Executor</code> releases</u> :",id:"3-removing-dependency-between-plugin-and-executor-releases-",children:[]}]},{value:"Result:",id:"result",children:[]},{value:"Other Considerations:",id:"other-considerations",children:[]}],c={toc:p};function d(e){var n=e.components,t=(0,i.Z)(e,l);return(0,r.kt)("wrapper",(0,a.Z)({},c,t,{components:n,mdxType:"MDXLayout"}),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Feature Name: Simplify Plugin Releases"),(0,r.kt)("li",{parentName:"ul"},"Status: Draft"),(0,r.kt)("li",{parentName:"ul"},"Start Date: 2022-05-07"),(0,r.kt)("li",{parentName:"ul"},"Author: Saikumar")),(0,r.kt)("h1",{id:"summary"},"Summary"),(0,r.kt)("p",null,"The scope of this rfc is to simplify the release and deployment operations w.r.t the optimus plugin ecosystem."),(0,r.kt)("p",null,"The proposal here is to remove :"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("strong",{parentName:"li"},"Executor and Task Dependency")," :",(0,r.kt)("br",{parentName:"li"}),"Decouple the ",(0,r.kt)("em",{parentName:"li"},"executor_boot_process")," and the executor as separate containers where the airflow worker launches a pod with ",(0,r.kt)("em",{parentName:"li"},"init-container")," (for boot process) adjacent to executor container."),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("strong",{parentName:"li"},"Plugin and Server Dependency")," :",(0,r.kt)("br",{parentName:"li"}),"Install plugins on-demand declaratively instead of manually baking them into the optimus server image (in kubernetes setup)."),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("strong",{parentName:"li"},"Plugin and Executor Dependency")," :",(0,r.kt)("br",{parentName:"li"}),"Extract out executor related variables from plugin code (Executor version , Image etc..) as ",(0,r.kt)("em",{parentName:"li"},"plugin config"),".")),(0,r.kt)("h1",{id:"technical-design"},"Technical Design"),(0,r.kt)("h2",{id:"background-"},"Background :"),(0,r.kt)("h3",{id:"optimus-components-that-trigger-a-new-release"},"Optimus components that trigger a new release:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Executor Image "),(0,r.kt)("li",{parentName:"ul"},"Task Image"),(0,r.kt)("li",{parentName:"ul"},"Plugin binary"),(0,r.kt)("li",{parentName:"ul"},"Optimus server/cli binary")),(0,r.kt)("h3",{id:"non-trivial-release-dependencies-as-per-current-design"},"Non-trivial release dependencies as per current design"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"Executor Image release")," -> ",(0,r.kt)("inlineCode",{parentName:"li"},"Task Image release")," -> ",(0,r.kt)("inlineCode",{parentName:"li"},"Plugin binary release")," -> ",(0,r.kt)("inlineCode",{parentName:"li"},"Server release")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"Plugin binary release")," -> ",(0,r.kt)("inlineCode",{parentName:"li"},"Server release"))),(0,r.kt)("h3",{id:"1-dependency-between-executor-and-task--"},"1. ",(0,r.kt)("u",null,"Dependency between ",(0,r.kt)("inlineCode",{parentName:"h3"},"Executor and Task")),"  :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"The ",(0,r.kt)("inlineCode",{parentName:"li"},"executor_boot_process")," and ",(0,r.kt)("inlineCode",{parentName:"li"},"executor")," are coupled:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"-- Plugin repo structure\n\n/task/\n/task/Dockerfile           -- task_image\n/task/executor/Dockerfile  -- executor_image\n")),(0,r.kt)("p",null,"Task Image :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"It's a wrapper around the executor_image to facilitate boot mechanism for executor."),(0,r.kt)("li",{parentName:"ul"},"The optimus binary is downloaded during buildtime of this image."),(0,r.kt)("li",{parentName:"ul"},"During runtime, it does as follow :",(0,r.kt)("ul",{parentName:"li"},(0,r.kt)("li",{parentName:"ul"},"Fetch assets, secrets, env from optimus server."),(0,r.kt)("li",{parentName:"ul"},"Load the env and launches the executor process.")))),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"task_image \n    | executor_image\n    | optimus-bin\n    | entrypoint.sh (load assets, env and launch executor)\n")),(0,r.kt)("p",null,"The ",(0,r.kt)("inlineCode",{parentName:"p"},"optimus-bin")," and ",(0,r.kt)("inlineCode",{parentName:"p"},"entrypoint.sh")," are baked into the ",(0,r.kt)("inlineCode",{parentName:"p"},"task_image")," and is being maintained by task/plugin developers."),(0,r.kt)("h3",{id:"2-dependency-between-plugin-and-server-also--plugin-and-executor-"},"2. ",(0,r.kt)("u",null,"Dependency between ",(0,r.kt)("inlineCode",{parentName:"h3"},"Plugin and Server"),", also  ",(0,r.kt)("inlineCode",{parentName:"h3"},"Plugin and Executor"))," :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Plugin binaries are manually installed (baked into optimus image in kubernetes setup)."),(0,r.kt)("li",{parentName:"ul"},"The executor_image and version are hard-coded into plugin binary. So, any change in executor version triggers additional release. (plugin-executor dependency)"),(0,r.kt)("li",{parentName:"ul"},"Any change in plugin code demands re-creation of optimus image with new plugin binary, inturn demanding redeployment of optimus server. (in kubernetes setup)")),(0,r.kt)("h2",{id:"approach-"},"Approach :"),(0,r.kt)("h3",{id:"1-removing-dependency-between-executor-and-task-image-releases-"},"1. ",(0,r.kt)("u",null,"Removing dependency between ",(0,r.kt)("inlineCode",{parentName:"h3"},"Executor and Task Image")," releases")," :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Decouple the lifecycle of the executor and the boot process as seperate containers/images.")),(0,r.kt)("img",{src:"images/simplify_plugins_executor.png",alt:"Simplify Plugins Executor",width:"800"}),(0,r.kt)("p",null,(0,r.kt)("strong",{parentName:"p"},"Task Boot Sequence"),":"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"KubernetesPodOperator spawns init-container and executor-container, mounted with shared volume (type emptyDir) for assets."),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("inlineCode",{parentName:"li"},"init-container")," fetches assets, config, env files and writes onto the shared volume."),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("inlineCode",{parentName:"li"},"postStart")," lifecycle hook in the ",(0,r.kt)("inlineCode",{parentName:"li"},"executor-container")," loads env from files on the shared volume."),(0,r.kt)("li",{parentName:"ol"},"the default entrypoint in the executor-image starts the actual job.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'# sample task definition\napiVersion: v1\nkind: Pod\nmetadata:\n  name: {{.task}}\nspec:\n  # init container\n  initContainers:\n  - name: init-executor\n    image: {{.default-init-docker-image}}\n    volumeMounts:\n    - mountPath: /usr/share/asserts\n      name: assets-dir\n  containers:\n    # executor container\n    - image: {{.executor-image-repo-link}}\n      name : {{.executor-name}}\n      volumeMounts:\n        - name: assets-dir\n          mountPath: /var/assets\n      # entrypoint.sh\n      lifecycle:\n        postStart:\n          exec:\n            command:\n                - "sh"\n                - "-c"\n                - >\n                source ~/.env\n                # more...\n\n  # shared volume\n  volumes:\n    - name: assets-dir\n      emptyDir: {}\n')),(0,r.kt)("h3",{id:"2-removing-dependency-between-plugin-and-server-releases-"},"2. ",(0,r.kt)("u",null,"Removing dependency between ",(0,r.kt)("inlineCode",{parentName:"h3"},"Plugin and Server")," releases")," :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Install plugin binaries on-demand with a declarative config (instead of baking them with optimus docker image - in kubernetes context)."),(0,r.kt)("li",{parentName:"ul"},"A plugin manager that consumes a declarative config to install the plugins on-demand is warrented."),(0,r.kt)("li",{parentName:"ul"},"Optimus support for plugin management as below.",(0,r.kt)("ul",{parentName:"li"},(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"optimus plugin install -c plugin_config.yaml --mode server/client")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"optimus plugin clean")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"optimus plugin list")))),(0,r.kt)("li",{parentName:"ul"},"This plugin manager can be taken advantage to sync plugins at client side as well, where in client mode the plugin manger pulls plugin_config from server and installs the plugins."),(0,r.kt)("li",{parentName:"ul"},"Support for different kinds of plugin repositories (like s3, gcs, url, local file system etc..) gives the added flexibility and options to distribute and install the plugin binaries in different ways."),(0,r.kt)("li",{parentName:"ul"},"Example for the plugin config: ")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'    plugins :\n      install_location : ""\n\n      # list the providers (plugin repository kind)\n      providers :\n      - type : url\n        name : internal_url_xyz_org\n        url: http://<internal_url>\n        auth: \n      \n      - type : gcs\n        name : private_gcs_backend_team\n        bucket: <bucket>\n        service_account : <base64_encoded_service_account>\n      \n      \n      # list the plugins and their config\n      plugins :\n      - provider : internal_url_xyz_org\n        output: <plugin-name>\n        path : <plugin_name>.tar.gz\n        config: {}  # plguin variables (eg: executor image etc..)\n      \n      - provider : private_gcs_backend_team\n        output: <plugin-name>\n        path : <plugin_name>.zip\n        config: {}\n      \n')),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"This feature demands that the current implementation of plugin discovery is revisited."),(0,r.kt)("li",{parentName:"ul"},"Plugin config can be mounted as configmap in the kubernetes setup.")),(0,r.kt)("h3",{id:"3-removing-dependency-between-plugin-and-executor-releases-"},"3. ",(0,r.kt)("u",null,"Removing dependency between ",(0,r.kt)("inlineCode",{parentName:"h3"},"Plugin and Executor")," releases")," :"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"As per the above mentioned sample plugin config, the values for the executor dependent variables in plugin binary can be infered from the plugin config itself."),(0,r.kt)("li",{parentName:"ul"},"This decouples the release dependency between plugin binary and executor.")),(0,r.kt)("h2",{id:"result"},"Result:"),(0,r.kt)("img",{src:"images/simplify_plugins.png",alt:"Simplify Plugins",width:"800"}),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Below are the implications of the proposed design, assuming the setup is in kubernetes :",(0,r.kt)("ul",{parentName:"li"},(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("strong",{parentName:"li"},"On Executor Release")," : Change plugin config in the configmap and restart the optimus pod."),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("strong",{parentName:"li"},"On Plugin Release")," : Change the plugin repository link in the plugin conf and restart the optimus pod."),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("strong",{parentName:"li"},"On Optimus Server release")," : Update the init-container for the ",(0,r.kt)("em",{parentName:"li"},"executor_boot_process"),". (not always)")))),(0,r.kt)("h2",{id:"other-considerations"},"Other Considerations:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"An assumption here is that the ",(0,r.kt)("em",{parentName:"li"},"executor_boot_process")," to remain same for all  executors."),(0,r.kt)("li",{parentName:"ul"},"One possible way to deal with the need for customised init process is to let plugins devs also provide ",(0,r.kt)("inlineCode",{parentName:"li"},"custom-init-image"),"\nalong with ",(0,r.kt)("inlineCode",{parentName:"li"},"executor-image")," which will fallback to a ",(0,r.kt)("inlineCode",{parentName:"li"},"default-init-image")," if not provided."),(0,r.kt)("li",{parentName:"ul"},"Supporting the ",(0,r.kt)("inlineCode",{parentName:"li"},"custom-init-image")," would require changes in plugin interfaces and rendering airflow dag."),(0,r.kt)("li",{parentName:"ul"},"Declarative plugin installation would affect the current implementation of plugin discovery.")))}d.isMDXComponent=!0}}]);