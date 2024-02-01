start:
	docker compose up -d

stop:
	docker compose down

rebuild:
	docker compose up --build -d app

logs:
	docker logs -f app