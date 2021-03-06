import fileList from "./filelist";

const filepath = require("./filepath");

import popup from "./popup";

import {i18nDOM} from "../i18n";

export function filesPopup (dir,t){
    const popups = new popup(t);

    return new Promise((resolve)=>{
        let fragment = document.createDocumentFragment();
        let popup = document.createElement("div");

        popup.setAttribute("class","modal active");
        popup.innerHTML = require("../html/files-popup.html");
        fragment.append(popup);

        i18nDOM(t,fragment);

        const closePopup = ()=>{
            document.body.removeChild(popup);
        };

        fragment.getElementById("cancel").onclick = ()=>{
            closePopup();
            resolve(false);
        };

        fragment.getElementById("modal-overlay").onclick = ()=>{
            closePopup();
            resolve(false);
        };

        const fileModalPath = fragment.getElementById("file-modal-path");

        const f = new fileList({
            path: dir,
            target: fragment.getElementById("attachment-modal-list"),
            onclickCallback: (file)=>{
                if (file.isDir){
                    f.navigate(filepath.join(f.path,file.name));
                    fileModalPath.innerText = f.path;
                }else{
                    switch(filepath.ext(file.name)){
                        case "md": case "html":
                            popups.alert(document.body,t("error"),t("errCannotAttachMDorHTML"));
                            break;
                        default:
                            closePopup();
                            resolve(filepath.join(f.path,file.name));
                    }
                }
            },
            actions: [
                {
                    icon: "fas fa-trash-alt",
                    tooltip: "Delete",
                    callback:file=>{
                        popups.confirm(document.body,t("confirmDelete"), `Delete ${file.name}?`)
                            .then(confirm=>{
                                if (confirm){
                                    if (file.isDir){
                                        fetch(filepath.join("/admin/api/list/",filepath.join(f.path,file.name)),{
                                            method: "DELETE",
                                        })
                                            .then(res=>{
                                                if (!res.ok){
                                                    alert("Error delete directory");
                                                }
                                                f.reload();
                                            })
                                    }else{
                                        fetch(filepath.join("/admin/api/blob/",filepath.join(f.path,file.name)),{
                                            method: "DELETE",
                                        })
                                            .then(res=>{
                                                if (!res.ok){
                                                    alert("Error rename file");
                                                }
                                                f.reload();
                                            })
                                    }
                                }
                            });
                    }
                }, {
                    icon: "fas fa-edit",
                    tooltip: "Rename",
                    callback:file=>{
                        popups.prompt(document.body,"Rename",`Rename ${file.name} to`)
                            .then(fn=>{
                                if (fn){
                                    if (file.isDir){
                                        fetch(filepath.join("/admin/api/list/",filepath.join(f.path,file.name)),{
                                            method: "PUT",
                                            body: JSON.stringify(filepath.join(f.path,fn))
                                        })
                                            .then(res=>{
                                                if (!res.ok){
                                                    alert("Error rename directory");
                                                }
                                                f.reload();
                                            })
                                    }else{
                                        fetch(filepath.join("/admin/api/blob/",filepath.join(f.path,file.name)),{
                                            method: "PUT",
                                            body: JSON.stringify(filepath.join(f.path,fn))
                                        })
                                            .then(res=>{
                                                if (!res.ok){
                                                    alert("Error rename file");
                                                }
                                                f.reload();
                                            })
                                    }
                                }

                            });
                    }
                }]
        },t);

        fragment.getElementById("upload-file").onclick = ()=>{
            popups.upload(document.body,filepath.join("/admin/api/blob",f.path))
                .then(()=>f.reload());
        };

        fileModalPath.innerText = f.path;

        document.body.appendChild(popup);
    })
}

export default filesPopup;