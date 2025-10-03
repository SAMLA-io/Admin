# Admin API

A general admin API for all SAMLA services.

## Installation

1. Clone the repository

```bash
git clone https://github.com/samla-io/samla-admin.git
cd samla-admin
```

2. Create a `.env` file

```bash
CLERK_SECRET_KEY=your_clerk_secret_key
AUTHORIZED_ORG_ID=your_authorized_org_id
```

3. Build and run the Docker container

```bash
docker build -t samla-admin .
docker run -p 8080:8080 samla-admin
```

The API will be available at `http://localhost:8080`.

## Endpoints

### Organizations

#### Get /organizations/all

Returns all organizations.

**Headers:**
- Authorization: Bearer <JWT Token>

**Response:**

```json
{
    "data": [
        {
            "object": "organization",
            "id": "org_123",
            "name": "Organization 1",
            "slug": "Organization-1",
            "image_url": "",
            "has_image": true,
            "max_allowed_memberships": 1,
            "admin_delete_enabled": true,
            "public_metadata": {},
            "private_metadata": {},
            "created_by": "user_1234",
            "created_at": 1754520218386,
            "updated_at": 1754520219401,
        }
    ],
    "total_count": 1
}
```

#### Post /organizations/create

Creates a new organization.

**Headers:**
- Authorization: Bearer <JWT Token>

**Request Body:**

```json
{
    "name": "Organization 1",
    "slug": "Organization-1",
    "max_allowed_memberships": 1,
    "admin_delete_enabled": true,
    "public_metadata": {},
    "private_metadata": {}
}
```

**Response:**

```json
    {
        "object": "organization",
        "id": "org_123",
        "name": "Organization 1",
        "slug": "Organization-1",
        "image_url": "",
        "has_image": true,
        "max_allowed_memberships": 1,
        "admin_delete_enabled": true,
        "public_metadata": {},
        "private_metadata": {},
        "created_by": "user_1234",
        "created_at": 1754520218386,
        "updated_at": 1754520219401,
    }
```

#### PATCH /organizations/update

Updates an organization.

**Headers:**
- Authorization: Bearer <JWT Token>

**Query Parameters:**
- organization_id: The ID of the organization to update

**Request Body:**

```json
{
    "name": "Organization 1",
    "slug": "Organization-1",
    "max_allowed_memberships": 1,
    "admin_delete_enabled": true,
    "public_metadata": {},
    "private_metadata": {}
}
```

**Response:**

```json
    {
        "object": "organization",
        "id": "org_123",
        "name": "Organization 1",
        "slug": "Organization-1",
        "image_url": "",
        "has_image": true,
        "max_allowed_memberships": 1,
        "admin_delete_enabled": true,
        "public_metadata": {},
        "private_metadata": {},
        "created_by": "user_1234",
        "created_at": 1754520218386,
        "updated_at": 1754520219401,
    }
```

#### DELETE /organizations/delete

Deletes an organization.

**Query Parameters:**
- organization_id: The ID of the organization to delete

**Response:**

```json
{
    "message": "Organization org_123 deleted successfully"
}
```

### Users

#### Get /users/all
Returns all users.

**Headers:**
- Authorization: Bearer <JWT Token>

**Response:**

```json
{
    "data" : [
        {
            "object": "user",
            "id": "user_123",
            "first_name": "John",
            "last_name": "Doe",
            "email_addresses": [ 
                "email_address": "john.doe@example.com"
            ],
            "username": "john.doe",
            "profile_image_url": "",
            "has_image": true,
            "public_metadata": {},
            "private_metadata": {},
            "created_by": "user_1234",
            "created_at": 1754520218386,
            "updated_at": 1754520219401,
        }
    ],
    "total_count": 1
}
```
#### Get /users/get

Returns a user.

**Headers:**
- Authorization: Bearer <JWT Token>

**Query Parameters:**
- user_id: The ID of the user to get

**Response:**

```json
    {
        "object": "user",
        "id": "user_123",
        "first_name": "John",
        "last_name": "Doe",
        "email_addresses": [ 
            "email_address": "john.doe@example.com"
        ],
        "username": "john.doe",
        "profile_image_url": "",
        "has_image": true,
        "public_metadata": {},
        "private_metadata": {},
        "created_by": "user_1234",
        "created_at": 1754520218386,
        "updated_at": 1754520219401,
    }
```

#### POST /users/create

Creates a new user.

**Headers:**
- Authorization: Bearer <JWT Token>

**Request Body:**

```json
    {
        "username": "john.doe",
        "first_name": "John",
        "last_name": "Doe",
        "email_addresses": [ 
            "email_address": "john.doe@example.com"
        ],
        "public_metadata": {},
        "private_metadata": {},
    }
```

**Response:**

```json
    {
        "object": "user",
        "id": "user_123",
        "first_name": "John",
        "last_name": "Doe",
        "email_addresses": [ 
            "email_address": "john.doe@example.com"
        ],
        "username": "john.doe",
        "profile_image_url": "",
        "has_image": true,
        "public_metadata": {},
        "private_metadata": {},
        "created_by": "user_1234",
        "created_at": 1754520218386,
        "updated_at": 1754520219401,
    }
```

#### PATCH /users/update

Updates a user.

**Headers:**
- Authorization: Bearer <JWT Token>

**Query Parameters:**
- user_id: The ID of the user to update

**Request Body:**

```json
    {
        "username": "john.doe",
        "first_name": "John",
        "last_name": "Doe",
        "email_addresses": [ 
            "email_address": "john.doe@example.com"
        ],
        "public_metadata": {},
        "private_metadata": {},
    }
```

**Response:**

```json
    {
        "object": "user",
        "id": "user_123",
        "first_name": "John",
        "last_name": "Doe",
        "email_addresses": [ 
            "email_address": "john.doe@example.com"
        ],
        "username": "john.doe",
        "profile_image_url": "",
        "has_image": true,
        "public_metadata": {},
        "private_metadata": {},
        "created_by": "user_1234",
        "created_at": 1754520218386,
        "updated_at": 1754520219401,
    }
```

#### DELETE /users/delete

Deletes a user.

**Headers:**
- Authorization: Bearer <JWT Token>

**Query Parameters:**
- user_id: The ID of the user to delete

**Response:**

```json
    {
        "id": "user_123",
        "object": "user",
        "deleted": true
    }
```

### Invitations

#### Get /invitations/all

Returns all invitations.

**Headers:**
- Authorization: Bearer <JWT Token>

**Response:**

```json
{
    "invitations": [
        {
            "id": "invitation_123",
            "object": "invitation",
            "email_address": "john.doe@example.com",
            "status": "pending",
            "created_at": 1754520218386,
            "updated_at": 1754520219401,
        }
    ],
    "total_count": 1
}
```

#### POST /invitations/create

Creates a new invitation.

**Headers:**
- Authorization: Bearer <JWT Token>

**Query Parameters:**
- user_id: The ID of the user to invite. The backend will get the email address from the user.

**Response:**

```json
{
    "id": "invitation_123",
    "object": "invitation",
    "email_address": "john.doe@example.com",
    "status": "pending",
    "created_at": 1754520218386,
    "updated_at": 1754520219401,
}
```


## License

This project is licensed under the a proprietary license. See the [LICENSE](LICENSE) file for details.

## Contributors

- [@jpgtzg](https://github.com/jpgtzg)

