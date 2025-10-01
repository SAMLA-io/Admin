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

```bash
GET /organizations
```

Returns all organizations.

**Headers:**
- Authorization: Bearer <JWT Token>

**Response:**

```json
{
    "organizations": [
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
    ]
}
```



## License

This project is licensed under the a proprietary license. See the [LICENSE](LICENSE) file for details.

## Contributors

- [@jpgtzg](https://github.com/jpgtzg)

