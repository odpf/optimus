"use strict";(self.webpackChunkoptimus=self.webpackChunkoptimus||[]).push([[1299],{3905:function(e,t,r){r.d(t,{Zo:function(){return c},kt:function(){return m}});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var p=n.createContext({}),l=function(e){var t=n.useContext(p),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},c=function(e){var t=l(e.components);return n.createElement(p.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,i=e.originalType,p=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),d=l(r),m=a,y=d["".concat(p,".").concat(m)]||d[m]||u[m]||i;return r?n.createElement(y,o(o({ref:t},c),{},{components:r})):n.createElement(y,o({ref:t},c))}));function m(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=r.length,o=new Array(i);o[0]=d;var s={};for(var p in t)hasOwnProperty.call(t,p)&&(s[p]=t[p]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var l=2;l<i;l++)o[l]=r[l];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},8216:function(e,t,r){r.r(t),r.d(t,{frontMatter:function(){return s},contentTitle:function(){return p},metadata:function(){return l},toc:function(){return c},default:function(){return d}});var n=r(7462),a=r(3366),i=(r(7294),r(3905)),o=["components"],s={id:"create-bigquery-view",title:"Create bigquery view"},p=void 0,l={unversionedId:"guides/create-bigquery-view",id:"guides/create-bigquery-view",isDocsHomePage:!1,title:"Create bigquery view",description:"A view is a virtual table defined by a SQL query. When you create a view,",source:"@site/docs/guides/create-bigquery-view.md",sourceDirName:"guides",slug:"/guides/create-bigquery-view",permalink:"/optimus/docs/guides/create-bigquery-view",editUrl:"https://github.com/odpf/optimus/edit/master/docs/docs/guides/create-bigquery-view.md",tags:[],version:"current",lastUpdatedBy:"Dery Rahman Ahaddienata",lastUpdatedAt:1658372997,formattedLastUpdatedAt:"7/21/2022",frontMatter:{id:"create-bigquery-view",title:"Create bigquery view"},sidebar:"docsSidebar",previous:{title:"Create bigquery table",permalink:"/optimus/docs/guides/create-bigquery-table"},next:{title:"Create bigquery external table",permalink:"/optimus/docs/guides/create-bigquery-external-table"}},c=[{value:"Creating table with Optimus",id:"creating-table-with-optimus",children:[]},{value:"Creating table over REST",id:"creating-table-over-rest",children:[]},{value:"Creating table over GRPC",id:"creating-table-over-grpc",children:[]}],u={toc:c};function d(e){var t=e.components,r=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,n.Z)({},u,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("p",null,"A view is a virtual table defined by a SQL query. When you create a view,\nyou query it in the same way you query a table. When a user queries the view,\nthe query results contain data only from the tables and fields specified in the\nquery that defines the view.\nAt the moment only standard view is supported."),(0,i.kt)("p",null,"There are 3 ways to create a view:"),(0,i.kt)("h3",{id:"creating-table-with-optimus"},"Creating table with Optimus"),(0,i.kt)("p",null,"Supported datastore can be selected by calling"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"optimus resource create\n")),(0,i.kt)("p",null,"Optimus will request a resource name which should be unique across whole datastore.\nAll resource specification contains a name field which conforms to a fixed format.\nIn case of bigquery view, format should be\n",(0,i.kt)("inlineCode",{parentName:"p"},"projectname.datasetname.viewname"),".\nAfter the name is provided, ",(0,i.kt)("inlineCode",{parentName:"p"},"optimus")," will create a file in configured datastore\ndirectory. Open the created specification file and add additional spec details\nas follows:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},'version: 1\nname: temporary-project.optimus-playground.first_view\ntype: view\nlabels:\n  usage: testview\n  owner: optimus\nspec:\n  description: "example description"\n  view_query: |\n    Select * from temporary-project.optimus-playground.first_table\n')),(0,i.kt)("p",null,"This will add labels, description, along with the query for view once the\n",(0,i.kt)("inlineCode",{parentName:"p"},"deploy")," command is invoked.\nTo use text editor intellisense for SQL formatting and linting, view query can\nalso be added in a separate file inside the same directory with the name ",(0,i.kt)("inlineCode",{parentName:"p"},"view.sql"),".\nDirectory will look something like:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"./\n./bigquery/temporary-project.optimus-playground.first_view/resource.yaml\n./bigquery/temporary-project.optimus-playground.first_view/view.sql\n")),(0,i.kt)("p",null,"Remove the ",(0,i.kt)("inlineCode",{parentName:"p"},"view_query")," field from the resource specification if the query is\nspecified in a seperate file."),(0,i.kt)("h3",{id:"creating-table-over-rest"},"Creating table over REST"),(0,i.kt)("p",null,"Optimus exposes Create/Update rest APIS"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"Create: POST /api/v1beta1/project/{project_name}/namespace/{namespace}/datastore/{datastore_name}/resource\nUpdate: PUT /api/v1beta1/project/{project_name}/namespace/{namespace}/datastore/{datastore_name}/resource\nRead: GET /api/v1beta1/project/{project_name}/namespace/{namespace}/datastore/{datastore_name}/resource/{resource_name}\n")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-json"},'{\n  "resource": {\n    "version": 1,\n    "name": "temporary-project.optimus-playground.first_view",\n    "datastore": "bigquery",\n    "type": "view",\n    "labels": {\n      "usage": "testview",\n      "owner": "optimus"\n    },\n    "spec": {\n      "description": "example description",\n      "view_query": "Select * from temporary-project.optimus-playground.first_table"\n    }\n  }\n}\n')),(0,i.kt)("h3",{id:"creating-table-over-grpc"},"Creating table over GRPC"),(0,i.kt)("p",null,"Optimus in RuntimeService exposes an RPC "),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-protobuf"},"rpc CreateResource(CreateResourceRequest) returns (CreateResourceResponse) {}\n\nmessage CreateResourceRequest {\n    string project_name = 1;\n    string datastore_name = 2;\n    ResourceSpecification resource = 3;\n    string namespace = 4;\n}\n")),(0,i.kt)("p",null,"Function payload should be self-explanatory other than the struct ",(0,i.kt)("inlineCode",{parentName:"p"},"spec")," part which\nis very similar to how json representation look."),(0,i.kt)("p",null,"Spec will use ",(0,i.kt)("inlineCode",{parentName:"p"},"structpb")," struct created with ",(0,i.kt)("inlineCode",{parentName:"p"},"map[string]interface{}"),"\nFor example:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'map[string]interface{\n    "description": "example description",\n    "view_query": "Select * from temporary-project.optimus-playground.first_table"\n}\n')))}d.isMDXComponent=!0}}]);