5. Executar o Docker Compose
   Navegue até o diretório backend/ e execute o Docker Compose:

bash
Copy code
docker-compose up
Este comando fará o seguinte:

Iniciará o contêiner MySQL.
Aplicará automaticamente as migrações usando o contêiner golang-migrate.
Iniciará o aplicativo Go após o banco de dados estar pronto.
6. Verificar as Migrações
   Após executar o comando acima, as tabelas devem estar criadas no banco de dados MySQL. Você pode conectar ao MySQL para verificar:

bash
Copy code
docker exec -it mysql mysql -uuser -ppassword face_door_admin_panel
No console MySQL, você pode listar as tabelas:

sql
Copy code
SHOW TABLES;
Você verá as tabelas users, events, e notifications se as migrações foram aplicadas corretamente.