/*!
 * HugoCMS
 * https://github.com/dfkdream/HugoCMS
 *
 * Copyright 2020 dfkdream
 * Released under MIT License
 */
const filepath = require('./filepath');

import fileList from "./filelist";

import popups from "./popup";

require('../css/style.css');
require('../css/list.css');

require('spectre.css/dist/spectre.min.css');

require('../css/all.min.css');

import publish from "./publish";

function fileToEndpoint(file){
   if (file.isDir) return "/admin/list";
   switch (filepath.ext(file.name)){
       case "md": case "html":
           return "/admin/edit";
       default:
           return "/admin/api/blob";
   }
}

import {i18n} from "../i18n";

i18n().then(t=> {
    const popup = new popups(t);
    const path = location.pathname.replace(/^(\/admin\/list)/, "");

    const locationHeader = document.getElementById("location");
    locationHeader.innerText = path;

    const f = new fileList({
        path: path,
        target: document.getElementById("list-tbody"),
        onclickCallback: (file) => {
            if (file.isDir) {
                f.navigate(filepath.join(f.path, file.name));
                locationHeader.innerText = f.path;
                history.pushState(f.path, "HugoCMS - " + f.path, filepath.join("/admin/list", f.path));
                document.title = "HugoCMS - " + f.path;
            } else {
                location.href = filepath.join(filepath.join(fileToEndpoint(file), f.path), file.name);
            }
        },
        actions: [
            {
                icon: "fas fa-trash-alt",
                tooltip: "Delete",
                callback: file => {
                    popup.confirm(document.body, t("confirmDelete"), `Delete ${file.name}?`)
                        .then(confirm => {
                            if (confirm) {
                                if (file.isDir) {
                                    fetch(filepath.join("/admin/api/list/", filepath.join(f.path, file.name)), {
                                        method: "DELETE",
                                    })
                                        .then(res => {
                                            if (!res.ok) {
                                                alert("Error delete directory");
                                            }
                                            f.reload();
                                        })
                                } else {
                                    fetch(filepath.join("/admin/api/blob/", filepath.join(f.path, file.name)), {
                                        method: "DELETE",
                                    })
                                        .then(res => {
                                            if (!res.ok) {
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
                callback: file => {
                    popup.prompt(document.body, t("rename"), `Rename ${file.name} to`)
                        .then(fn => {
                            if (fn) {
                                if (file.isDir) {
                                    fetch(filepath.join("/admin/api/list/", filepath.join(f.path, file.name)), {
                                        method: "PUT",
                                        body: JSON.stringify(filepath.join(f.path, fn))
                                    })
                                        .then(res => {
                                            if (!res.ok) {
                                                alert("Error rename directory");
                                            }
                                            f.reload();
                                        })
                                } else {
                                    fetch(filepath.join("/admin/api/blob/", filepath.join(f.path, file.name)), {
                                        method: "PUT",
                                        body: JSON.stringify(filepath.join(f.path, fn))
                                    })
                                        .then(res => {
                                            if (!res.ok) {
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

    document.getElementById("new-directory").onclick = () => {
        popup.prompt(document.body, t("createDirectory"), t("enterDirectoryName"))
            .then(inp => {
                if (!inp) return;
                fetch(filepath.join(filepath.join("/admin/api/list", f.path), inp), {method: "POST"})
                    .then(() => f.reload());
            });
    };

    document.getElementById("new-post").onclick = () => {
        popup.prompt(document.body, t("newPost"), t("enterPostFilename"))
            .then(inp => {
                if (!inp) return;
                location.href = filepath.join(filepath.join("/admin/edit", f.path), inp);
            });
    };

    document.getElementById("rebuild").onclick = () => {
        publish(t);
    };

    document.getElementById("upload-file").onclick = () => {
        popup.upload(document.body, filepath.join("/admin/api/blob", f.path))
            .then(() => f.reload());
    };

    history.replaceState(f.path, "HugoCMS - " + f.path, filepath.join("/admin/list", f.path));
    document.title = "HugoCMS - " + f.path;

    window.onpopstate = (e) => {
        f.navigate(e.state);
        locationHeader.innerText = f.path;
    };
});