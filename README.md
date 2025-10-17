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

#### GET /organizations/users

Retrieves all users for an organization

**Query Parameters:**
- organization_id: The ID of the organization to delete

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
    "total_count" : 1
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

### Sarah

All endpoints for Sarah are under the `/sarah` endpoint.s

#### Organizations

##### GET /assistants/organizations

Retrieve all assistants for an organization.

**Headers:**
- `Authorization: Bearer <clerk_jwt_token>` (required)

**Query Parameters:**
- organization_id: The ID of the organization to delete

**Response:**
```json
[
  {
    "id": "507f1f77bcf86cd799439011",
    "name": "Insurance Reminder Assistant",
    "vapi_assistant_id": "asst_1234567890abcdef",
    "type": "insurance"
  }
]
```

#### POST /assistants/create

Create a new assistant.

**Headers:**
- `Authorization: Bearer <clerk_jwt_token>` (required)

**Request Body:**

Note: Not all fields are required to be passed. This object is the same as the VapiAI Assistant Object, which is also the same in the update endpoint.

```json
{
  "assistantUpdateRequest": {
    "transcriber": {
        "model": "nova-2",
        "language": "es",
        "numerals": false,
        "confidenceThreshold": 0.4,
        "endpointing": 300,
        "provider": "deepgram"
    },
    "model": {
        "messages": [
            {
                "content": "",
                "role": "system"
            }
        ],
        "model": "gpt-5",
        "temperature": 1,
        "provider": "openai"
    },
    "voice": {
        "voiceId": "",
        "stability": 0.5,
        "similarityBoost": 0.75,
        "model": "",
        "provider": ""
    },
    "firstMessage": "",
    "clientMessages": [
        "conversation-update",
        "function-call",
        "hang",
        "model-output",
        "speech-update",
        "status-update",
        "transfer-update",
        "transcript",
        "tool-calls",
        "user-interrupted",
        "voice-input",
        "workflow.node.started"
    ],
    "serverMessages": [
        "conversation-update",
        "end-of-call-report",
        "function-call",
        "hang",
        "speech-update",
        "status-update",
        "tool-calls",
        "transfer-destination-request",
        "user-interrupted"
    ],
    "backgroundDenoisingEnabled": true,
    "name": "",
    "voicemailMessage": "",
    "endCallMessage": "",
    "analysisPlan": {
        "summaryPlan": {
            "messages": [
                {
                    "content": "",
                    "role": "system"
                },
                {
                    "content": "",
                    "role": "user"
                }
            ]
        },
        "structuredDataPlan": {
            "messages": [
                {
                    "content": "",
                    "role": "system"
                },
                {
                    "content": "",
                    "role": "user"
                }
            ],
            "enabled": true,
            "schema": {
                "type": "object",
                "properties": {
                    "Categoria": {
                        "type": "string"
                    },
                    "name": {
                        "type": "string"
                    },
                    "phonenumber": {
                        "type": "string"
                    }
                },
                "required": [
                    "name",
                    "Categoria"
                ]
            }
        },
        "successEvaluationPlan": {
            "rubric": "PercentageScale",
            "messages": [
                {
                    "content": "",
                    "role": "system"
                },
                {
                    "content": "",
                    "role": "user"
                }
            ]
        }
    },
    "artifactPlan": {
        "recordingFormat": "mp3"
    },
    "startSpeakingPlan": {
        "waitSeconds": 0.4,
        "smartEndpointingEnabled": "livekit",
        "transcriptionEndpointingPlan": {
            "onPunctuationSeconds": 0.1,
            "onNoPunctuationSeconds": 1.5,
            "onNumberSeconds": 0.5
        }
    },
    "id": "689eadfb-5f57-4b45-a7b1-7b37fa5baaf8",
    "orgId": "7331923f-e36c-4bb9-938d-9a6c651a2f49",
    "createdAt": "2025-09-05T17:10:33Z",
    "updatedAt": "2025-09-25T20:14:21Z"    
  }
}
```

**Response:**

```json
{
  "InsertedID": "507f1f77bcf86cd799439011",
  "Acknowledged": true
}
```

#### POST /assistants/register
Register an assistant in the database. This is useful when an assistant is already created in VapiAI and needs to be registered in the database manually.

This endpoint will not create the assistant in VapiAI, it will only register the assistant in the database IF it exists in VapiAI.

**Headers:**
- `Authorization: Bearer <clerk_jwt_token>` (required)

**Request Body:**

```json
{
  "assistant": {
    "name": "foo",
    "vapiAssistantId": "foo",
    "type": "foo"
  }
}
```

**Response:**

```json
{
  "InsertedID": "507f1f77bcf86cd799439011",
  "Acknowledged": true
}
```

#### DELETE /assistants/delete
Deletes an assistant.

**Headers:**
- `Authorization: Bearer <clerk_jwt_token>` (required)

**Query Parameters:**
- `assistantId`: The VapiAI assistant ID to delete (required)

**Response:**

```json
{
  "acknowledged": true,
  "deletedCount": 1
}
```


## License

This project is licensed under the a proprietary license. See the [LICENSE](LICENSE) file for details.

## Contributors

- [@jpgtzg](https://github.com/jpgtzg)

