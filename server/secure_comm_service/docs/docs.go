// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/files/all": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Возвращает пре‐подписанные ссылки на скачивание всех файлов заданной категории (photo, unknown, video, text).",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Получение всех файлов категории",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Категория файлов (photo, unknown, video, text)",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список ссылок на все файлы категории",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.FileResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Некорректная категория",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "Доступ запрещён",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Файлы не найдены",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/files/many": {
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Удаляет несколько объектов из MinIO, переданных в JSON-массиве, и снижает использование квоты.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Удаление нескольких файлов",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Массив идентификаторов объектов",
                        "name": "objectIDs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ObjectIDs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Файлы успешно удалены",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Некорректный JSON в теле запроса",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Доступ запрещён",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Один из файлов не найден",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/files/one": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Возвращает пре‐подписанную ссылку на скачивание одного файла по ID и типу.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Получение одного файла",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Идентификатор объекта",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Категория файла (photo, unknown, video, text)",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ссылка на скачивание файла",
                        "schema": {
                            "$ref": "#/definitions/dto.FileResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Доступ запрещён",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Файл не найден",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Удаляет один объект из MinIO и снижает использование квоты.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Удаление одного файла",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Идентификатор объекта",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Категория файла (photo, unknown, video, text)",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Файл успешно удалён",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Доступ запрещён",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Файл не найден",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/cloud_handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/files/one/encrypted": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Токен авторизации + зашифрованный поток в body + метаданные в заголовках",
                "consumes": [
                    "application/octet-stream"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Загрузка зашифрованного файла в MinIO(объектное хранилище) “на лету”",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Оригинальное имя файла (например photo.jpg)",
                        "name": "X-Orig-Filename",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Оригинальный MIME-тип (например image/jpeg)",
                        "name": "X-Orig-Mime",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Категория файла (photo, video, text, unknown)",
                        "name": "X-File-Category",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Зашифрованный бинарный поток (application/octet-stream)",
                        "name": "file",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешно: JSON с metadata + presigned URL",
                        "schema": {
                            "$ref": "#/definitions/dto.FileResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректные заголовки или body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Проблема с авторизацией",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/handshake/finalize": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "ЗАПРОС ОТ КЛИЕНТА:\nКлиент шлёт RSA-OAEP(encrypted payload), закодированный в Base64.\nи signature подписанный payload приватным ключем клиента\nоб отправляемых полях клиентом encrypted и signature3:\nРандомные 32 байта - это сессионная строка, назовем её ks, которая лежит в payload\npayload - это сумма байтов (ks || nonce3 || nonce2)\nsignature3 - это подписанный payload приватным ключем ECDSA клиента в base64\n\nОТВЕТ ОТ СЕРВЕРА:\nСервер возвращает подпись signature4 = SHA256(Ks || nonce3 || nonce2), подписанную приватным ECDSA-ключом сервера и закодированную в Base64.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "handshake"
                ],
                "summary": "Завершает Handshake",
                "parameters": [
                    {
                        "description": "Параметры завершения Handshake",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.HandshakeFinalizeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.HandshakeFinalizeResp"
                        }
                    },
                    "400": {
                        "description": "Некорректный JSON или параметры",
                        "schema": {
                            "$ref": "#/definitions/dto.BadRequestErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized или подпись не верна",
                        "schema": {
                            "$ref": "#/definitions/dto.UnauthorizedErr"
                        }
                    },
                    "409": {
                        "description": "Conflict — повторный запрос (replay-detected)",
                        "schema": {
                            "$ref": "#/definitions/dto.ConflictErr"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/handshake/init": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "ЗАПРОС ОТ КЛИЕНТА:\nКлиент отправляет свои публичные ключи и nonce1, всё это подписано приватным ECDSA-ключом.\nВсе бинарные данные (ключи, подписи, nonce) закодированы в Base64 (DER для ключей и подписи).\n\nrsa_pub_client - Base64(DER-закодированный RSA-публичный ключ клиента)\necdsa_pub_client - Base64(DER-закодированный ECDSA-публичный ключ клиента)\nnonce1 - Base64(8-байтовый случайный nonce)\nsignature1 - Base64(DER-закодированная подпись SHA256(clientRSA || clientECDSA || nonce1) приватным ECDSA-ключом клиента)\n\nОТВЕТ ОТ СЕРВЕРА:\nСервер отвечает своими публичными ключами и nonce2, всё это подписано приватным ECDSA-ключом сервера.\nВсе бинарные данные (ключи, подписи, nonce) закодированы в Base64 (DER для ключей и подписи).\n\nclient_id будет извлечён из JWT (Bearer-токена).\nrsa_pub_server - Base64(DER-закодированный RSA-публичный ключ сервера)\necdsa_pub_server - Base64(DER-закодированный ECDSA-публичный ключ сервера)\nnonce2 - Base64(8-байтовый случайный nonce)\nsignature2 - Base64(DER-подпись SHA256(rsaServer || ecdsaServer || nonce2 || nonce1 || clientID) приватным ECDSA-ключом сервера)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "handshake"
                ],
                "summary": "Инициализация Handshake",
                "parameters": [
                    {
                        "description": "Параметры инициации Handshake",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.HandshakeInitReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.HandshakeInitResp"
                        }
                    },
                    "400": {
                        "description": "Некорректный JSON или параметры",
                        "schema": {
                            "$ref": "#/definitions/dto.BadRequestErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized или ошибка подписи",
                        "schema": {
                            "$ref": "#/definitions/dto.UnauthorizedErr"
                        }
                    },
                    "409": {
                        "description": "Conflict — повторный запрос (replay-detected)",
                        "schema": {
                            "$ref": "#/definitions/dto.ConflictErr"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/session/test": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Клиент шлёт серверу зашифрованный пакет, объединяющий метаданные и payload данные:\ntimestamp - время отправки клиента в формате Unix миллисекунд (8 байт)\nnonce - криптографически стойкое случайное значение (16 байт) для защиты от replay\nIV - инициализационный вектор AES-CBC (16 байт)\nciphertext - результат AES-256-CBC шифрования сессионного payload'а + PKCS#7\ntag - HMAC-SHA256 от (IV||ciphertext) для целостности данных\n\nпо итогу клиент отправляет:\nencrypted_message = timestamp(8 byte) || nonce(16 byte) || IV(16 byte) || ciphertext || tag(32 byte)\nclient_signature = Base64(DER‑закодированная подпись SHA256(timestamp || nonce || IV || ciphertext || tag) приватным ECDSA‑ключом клиента\n\nСервер выполняет следующие шаги:\nДекодирует Base64 и извлекает blob.\nИзвлекает timestamp и nonce из первых 24 байт и проверяет не использован ли nonce и если diff = (client_timestamp - server_timestamp) \u003e допустимому окну(30секунд). если это так, то запрос отклоняется с ошибкой\nПолучает K_enc и K_mac по clientID, иначе 401 Unauthorized;\nПроверяет HMAC-SHA256(iv||ciphertext), иначе 401 Unauthorized;\nРасшифровывает AES-256-CBC, снимает PKCS#7 padding, иначе 400 Bad Request;\nВозвращает JSON с plaintext - декодированный userData.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "session"
                ],
                "summary": "Тестовое расшифрование и проверка целостности сессионного сообщения",
                "parameters": [
                    {
                        "description": "Метаданные + зашифрованный payload в Base64",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SessionMessageReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ: plaintext",
                        "schema": {
                            "$ref": "#/definitions/dto.SessionMessageResp"
                        }
                    },
                    "400": {
                        "description": "Неверный формат Base64, устаревший timestamp или padding",
                        "schema": {
                            "$ref": "#/definitions/dto.BadRequestErr"
                        }
                    },
                    "401": {
                        "description": "Session not found или проверка HMAC не прошла",
                        "schema": {
                            "$ref": "#/definitions/dto.UnauthorizedErr"
                        }
                    },
                    "409": {
                        "description": "Повторное использование nonce (replay)",
                        "schema": {
                            "$ref": "#/definitions/dto.ConflictErr"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/user/{id}/plan/init": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Создаёт для пользователя с переданным ID бесплатный план хранения объёмом до 10 GiB. Если план уже существует, операция игнорируется.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quota"
                ],
                "summary": "Инициализация бесплатного плана(10гб) пользователя. Вызывается как внешнее апи, при регистрации(/user/signup)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created — план успешно инициализирован",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Некорректный ID пользователя",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Ошибка авторизации или токен не предоставлен",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user/{id}/usage": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает, сколько гигабайт, мегабайт и килобайт хранится у пользователя, а также лимит (10 GiB для бесплатного плана) и имя плана.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quota"
                ],
                "summary": "Получение текущего использования диска пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Информация об использовании дискового пространства",
                        "schema": {
                            "$ref": "#/definitions/dto.UserUsage"
                        }
                    },
                    "400": {
                        "description": "Некорректный ID пользователя",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Ошибка авторизации или токен не предоставлен",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Пользователь не найден",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cloud_handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "description": "Дополнительные детали (при наличии)\nexample: \"field 'file' is required\""
                },
                "error": {
                    "description": "Краткое описание ошибки\nexample: \"Invalid request\"",
                    "type": "string"
                },
                "status": {
                    "description": "Код HTTP-статуса\nexample: 400",
                    "type": "integer"
                }
            }
        },
        "dto.BadRequestErr": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "example: invalid request data",
                    "type": "string"
                }
            }
        },
        "dto.ConflictErr": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "example: replay detected",
                    "type": "string"
                }
            }
        },
        "dto.FileResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "mime_type": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "obj_id": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "dto.HandshakeFinalizeReq": {
            "description": "Клиент шлёт RSA-OAEP(encrypted payload), закодированный в Base64. Подробнее про поле encrypted... Рандомные 32 байта - это сессионная строка, назовем её ks, которая лежит в payload payload - это сумма байтов (ks || nonce3 || nonce2) signature3 - это подписанный payload приватным ключем клиента В конце encrypted это зашифрованные байты (payload || signature3(в DER формате)) encrypted - зашифрован RSA-OAEP публичным ключем сервера, отдается в формате Base64",
            "type": "object",
            "properties": {
                "encrypted": {
                    "description": "Base64(RSA-OAEP(encrypted payload || signature3(DER)))",
                    "type": "string"
                },
                "signature3": {
                    "type": "string"
                }
            }
        },
        "dto.HandshakeFinalizeResp": {
            "description": "Сервер возвращает подпись h4 = SHA256(Ks || nonce3 || nonce2), подписанную приватным ECDSA‑ключом сервера и закодированную в Base64.",
            "type": "object",
            "properties": {
                "signature4": {
                    "description": "Base64(DER‑подпись ответа сервера)",
                    "type": "string"
                }
            }
        },
        "dto.HandshakeInitReq": {
            "description": "Клиент отправляет свои публичные ключи и nonce1, всё это подписано приватным ECDSA‑ключом. Все бинарные данные (ключи, подписи, nonce) закодированы в Base64 (DER для ключей и подписи).  rsa_pub_client - Base64(DER‑закодированный RSA‑публичный ключ клиента) ecdsa_pub_client - Base64(DER‑закодированный ECDSA‑публичный ключ клиента) nonce1 - Base64(8‑байтовый случайный nonce) signature1 - Base64(DER‑закодированная подпись SHA256(clientRSA || clientECDSA || nonce1) приватным ECDSA‑ключом клиента)",
            "type": "object",
            "properties": {
                "ecdsa_pub_client": {
                    "type": "string"
                },
                "nonce1": {
                    "type": "string"
                },
                "rsa_pub_client": {
                    "type": "string"
                },
                "signature1": {
                    "type": "string"
                }
            }
        },
        "dto.HandshakeInitResp": {
            "description": "Сервер отвечает своими публичными ключами и nonce2, всё это подписано приватным ECDSA‑ключом сервера. Все бинарные данные (ключи, подписи, nonce) закодированы в Base64 (DER для ключей и подписи).  client_id - SHA256‑хэш от (clientRSA‖clientECDSA), представлен в hex rsa_pub_server - Base64(DER‑закодированный RSA‑публичный ключ сервера) ecdsa_pub_server - Base64(DER‑закодированный ECDSA‑публичный ключ сервера) nonce2 - Base64(8‑байтовый случайный nonce) signature2 - Base64(DER‑подпись SHA256(rsaServer || ecdsaServer || nonce2 || nonce1 || clientID) приватным ECDSA‑ключом сервера)",
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "ecdsa_pub_server": {
                    "type": "string"
                },
                "nonce2": {
                    "type": "string"
                },
                "rsa_pub_server": {
                    "type": "string"
                },
                "signature2": {
                    "type": "string"
                }
            }
        },
        "dto.InternalServerErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dto.ObjectIDs": {
            "type": "object",
            "properties": {
                "object_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.SessionMessageReq": {
            "type": "object",
            "properties": {
                "client_signature": {
                    "type": "string"
                },
                "encrypted_message": {
                    "description": "Base64(IV || ciphertext || tag)",
                    "type": "string"
                }
            }
        },
        "dto.SessionMessageResp": {
            "type": "object",
            "properties": {
                "plaintext": {
                    "type": "string"
                }
            }
        },
        "dto.UnauthorizedErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dto.UserUsage": {
            "type": "object",
            "properties": {
                "current_used_gb": {
                    "type": "integer"
                },
                "current_used_kb": {
                    "type": "integer"
                },
                "current_used_mb": {
                    "type": "integer"
                },
                "plan_name": {
                    "type": "string"
                },
                "storage_limit_gb": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "\"Bearer {token}\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "SecureComm API",
	Description:      "Документация о внутренней реализации и логики работы находится в папке docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
