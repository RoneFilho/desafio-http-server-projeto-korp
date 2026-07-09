# Readme Desafio-http-server-projeto-korp


Primeiramente faça o clone do repositório e abra o terminal linux na pasta raiz do projeto, onde temos o compose.yaml e o playbook.yml

```bash
git clone https://github.com/RoneFilho/desafio-http-server-projeto-korp
```

Para rodar o Ansible playbook.yml tenha certeza que possui o Ansible instalado na máquina e basta executar o seguinte comando:

```bash
ansible-playbook playbook.yml
```

O playbook cuidará da instalação do docker e docker compose, subirá todos os containers (Http-server, Nginx, Prometheus e Grafana), irá configurar automaticamente o Grafana com o Datasource e Dashboard. 
No final irá retornar o conteúdo JSON da requisição Curl feita para http://localhost/projeto-korp

As aplicações estão nas seguintes portas:
localhost:80/project-korp (http-server)
localhost:9090 (Prometheus)
localhost:3000 (Grafana)

Para acessar o grafana basta usar: 
login: admin 
senha: admin

## Contato:

Atualmente estou aberto para oportunidades como DevOps e Desenvolvedor Backend focado em Python, C# e Java!

📬 Entre em contato:
📧 ronefilho28@gmail.com
💼 [Linkedin](https://www.linkedin.com/in/rone-pereira-alves-filho-401535224/)