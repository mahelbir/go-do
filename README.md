## Go-DO

Go-Do is a TODO application built with Golang. This API allows users to manage their TODO lists and tasks efficiently. The collection is structured into several main sections, each focusing on a different aspect of the application.

## Authentication

The API uses JWT Bearer tokens for authentication. The tokens for Admin and User roles are required to access different endpoints.

- **Admin Token**: `{{auth-admin}}`
- **User Token**: `{{auth-user}}`

## Base URL

All endpoints use the following base URL:
```
http://localhost:8080
```

## Endpoints

### Auth

#### Login Admin

- **Method**: POST
- **URL**: `{{base-url}}/auth/login`
- **Description**: Login as Admin
- **Request Body**:
  ```json
  {
      "email": "admin@gmail.com",
      "password": "admin123"
  }
  ```

#### Login User

- **Method**: POST
- **URL**: `{{base-url}}/auth/login`
- **Description**: Login as Regular User
- **Request Body**:
  ```json
  {
      "email": "user@gmail.com",
      "password": "user123"
  }
  ```

### User

#### Get User

- **Method**: GET
- **URL**: `{{base-url}}/users/1`
- **Description**: Access: Admin
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-admin}}"
  }
  ```

### TodoList

#### Create TodoList

- **Method**: POST
- **URL**: `{{base-url}}/todo_list`
- **Description**: Access: Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```
- **Request Body**:
  ```json
  {
      "title": "Test Title"
  }
  ```

#### List Owned TodoLists

- **Method**: GET
- **URL**: `{{base-url}}/todo_list`
- **Description**: Access: Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```

#### List All TodoLists

- **Method**: GET
- **URL**: `{{base-url}}/todo_list?all=true`
- **Description**: Access: Admin
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-admin}}"
  }
  ```

#### Get TodoList

- **Method**: GET
- **URL**: `{{base-url}}/todo_list/1`
- **Description**: Access: Admin/Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```

#### Update TodoList

- **Method**: PATCH
- **URL**: `{{base-url}}/todo_list/1`
- **Description**: Access: Admin/Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```
- **Request Body**:
  ```json
  {
      "title": "Updated Title"
  }
  ```

#### Delete TodoList

- **Method**: DELETE
- **URL**: `{{base-url}}/todo_list/1`
- **Description**: Access: Admin/Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```

### TodoMessage

#### Create TodoMessage

- **Method**: POST
- **URL**: `{{base-url}}/todo_message/1`
- **Description**: Access: Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```
- **Request Body**:
  ```json
  {
      "content": "Test content"
  }
  ```

#### List TodoMessages by TodoList

- **Method**: GET
- **URL**: `{{base-url}}/todo_message/list/1`
- **Description**: Access: Admin/Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```

#### Get TodoMessage

- **Method**: GET
- **URL**: `{{base-url}}/todo_message/1`
- **Description**: Access: Admin/Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```

#### Update TodoMessage

- **Method**: PATCH
- **URL**: `{{base-url}}/todo_message/1`
- **Description**: Access: Admin/Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```
- **Request Body**:
  ```json
  {
      "content": "Updated content",
      "is_completed": true
  }
  ```

#### Delete TodoMessage

- **Method**: DELETE
- **URL**: `{{base-url}}/todo_message/1`
- **Description**: Access: Admin/Owner
- **Headers**:
  ```json
  {
      "Authorization": "Bearer {{auth-user}}"
  }
  ```

## Postman

You can test on [Postman](https://www.postman.com/grey-trinity-401646/workspace/api-docs/collection/10662426-e31a49f5-8df2-4126-b980-08a6f7188ee2).



## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
