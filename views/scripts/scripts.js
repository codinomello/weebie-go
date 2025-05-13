var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g = Object.create((typeof Iterator === "function" ? Iterator : Object).prototype);
    return g.next = verb(0), g["throw"] = verb(1), g["return"] = verb(2), typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var _this = this;
// URL do backend
var backendUrl = "http://localhost:8080";
// Função para cadastrar usuário
var registerForm = document.getElementById("registerForm");
var emailInput = document.getElementById("email");
var passwordInput = document.getElementById("password");
var nameInput = document.getElementById("name");
var registerMessage = document.getElementById("registerMessage");
registerForm.addEventListener("submit", function (e) { return __awaiter(_this, void 0, void 0, function () {
    var email, password, name, response, result, error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                e.preventDefault();
                email = emailInput.value;
                password = passwordInput.value;
                name = nameInput.value;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 4, , 5]);
                return [4 /*yield*/, fetch("".concat(backendUrl, "/signup"), {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify({ email: email, password: password, name: name }),
                    })];
            case 2:
                response = _a.sent();
                return [4 /*yield*/, response.json()];
            case 3:
                result = _a.sent();
                if (response.ok) {
                    registerMessage.textContent = "Usu\u00E1rio cadastrado com sucesso! UID: ".concat(result.uid);
                    registerMessage.style.color = "green";
                }
                else {
                    registerMessage.textContent = "Erro: ".concat(result.error || "Falha no cadastro");
                    registerMessage.style.color = "red";
                }
                return [3 /*break*/, 5];
            case 4:
                error_1 = _a.sent();
                registerMessage.textContent = "Erro ao conectar ao servidor";
                registerMessage.style.color = "red";
                return [3 /*break*/, 5];
            case 5: return [2 /*return*/];
        }
    });
}); });
// Função para cadastrar projeto
var projectForm = document.getElementById("project-form");
var ownerUidInput = document.getElementById("owner-uid");
var titleInput = document.getElementById("title");
var detailsInput = document.getElementById("details");
var groupInput = document.getElementById("group");
var yearInput = document.getElementById("year");
var courseInput = document.getElementById("course");
var odsInput = document.getElementById("ods");
var projectMessage = document.getElementById("project-message");
projectForm.addEventListener("submit", function (e) { return __awaiter(_this, void 0, void 0, function () {
    var project, response, result, error_2;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                e.preventDefault();
                project = {
                    owner_uid: ownerUidInput.value,
                    title: titleInput.value,
                    details: detailsInput.value,
                    group: groupInput.value,
                    year: parseInt(yearInput.value),
                    course: courseInput.value,
                    ods: odsInput.value.split(",").map(Number),
                };
                _a.label = 1;
            case 1:
                _a.trys.push([1, 4, , 5]);
                return [4 /*yield*/, fetch("".concat(backendUrl, "/project"), {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify(project),
                    })];
            case 2:
                response = _a.sent();
                return [4 /*yield*/, response.json()];
            case 3:
                result = _a.sent();
                if (response.ok) {
                    projectMessage.textContent = "Projeto cadastrado com sucesso! ID: ".concat(result.id);
                    projectMessage.style.color = "green";
                }
                else {
                    projectMessage.textContent = "Erro: ".concat(result.error || "Falha no cadastro");
                    projectMessage.style.color = "red";
                }
                return [3 /*break*/, 5];
            case 4:
                error_2 = _a.sent();
                projectMessage.textContent = "Erro ao conectar ao servidor";
                projectMessage.style.color = "red";
                return [3 /*break*/, 5];
            case 5: return [2 /*return*/];
        }
    });
}); });
