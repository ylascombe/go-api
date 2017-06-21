

# Objectifs de l'API
Ce projet contient le code qui permet de construire l'API en GO de la plateforme réactive.

Elle permettra à terme :
* recenser les feature team et leurs applications
* génerer le manifest permettant de construire un environnement complet
* se substituer aux applications_vars et clés SSH stockés jusqu'ici dans le dépot `init-poste-dev`


# Dev setup

## Pré-requis

Compilateur GO en version >= 1.8.1
Docker-compose en version >= 1.8.0

# Run setup

```shell
make run
```

# Build setup

TODO

# Contact

Contacter l'équipe CORE

# Known bugs

TODO

# Changelog

v0.1 : 
* CRUD sur les objets métiers `team`, `product`, `module`, `environment`, `artefact` 

# TODO

Dans ce readme : définir les codes retours


# Objets métiers manipulés dans l'API

**User** : 
Utilisateur de l'API ARC.

**Team** :
Feature team : est constitué de plusieurs utilisateurs

**Role** :
Rôle d'un utilisateur.
Par exemple, developpeur, product owner, testeur...

**Environment** :
Environnement sur lequel est déployé une instance de l'architecture réactive + les applications 
Exemple : l'environnement RECETTE, DEV, PRE_PROD, PROD

**Product** :
Application déployable sur un environnement. Elle est composée d'un ou plusieurs modules qui peuvent être de type Spark ,API, Schema Registry ou Template elasticsearch.

**Module** :
Sous ensemble d'un Product. C'est un élément constitué d'un artefact et d'une configuration.

**Artefact** :
Artefact qui constitue le binaire (ou le code) permettant d'installer un module. Il est le résultat d'une phase de build lorsqu'il s'agit d'un code compilé.


# Principes suivis

* RESTful API non stricte dans laquelle on s'autorise d'ajouter des routes facilitant l'utilisation

## Dans les URI

La description des ressources est effectuée via des noms concrets, pas de verbe.
Utilisation des verbes HTTP pour indiquer l'action : 
* GET : utilisé pour lire une collection ou un élement de la collection (si identifiant fourni)
* POST : utilisé pour créer une instance au sein d’une collection. L’identifiant de la ressource à créer ne doit pas être précisé. Le code retour n’est pas 200 mais 201.                                                                                                                                   L’URI et l’identifiant de la nouvelle ressource sont retournés dans le header “Location” de la réponse.
* PATCH : utilisé pour une mise à jour partielle d’une instance de la collection
* PUT : utilisé systématiquement pour réaliser une mise à jour totale d’une instance de la collection (tous les attributs sont remplacés et ceux qui sont non présents seront supprimés) **OU** création si l'identifiant est fourni
* DELETE : suppression d'un élément

-----------|---------------------|----------------------------------|--------------------------------------------|
Verbe HTTP | Correspondance CRUD | Collection : /orders             | Instance : /orders/{id}                    |
-----------|---------------------|----------------------------------|--------------------------------------------|
GET 	   | READ 	             | Read a list orders. 200 OK. 	    | Read the detail of a single order. 200 OK. |
POST 	   | CREATE 	         | Create a new order. 201 Created. |                                            |
PUT 	   | UPDATE/CREATE 	 	 | Full Update. 200 OK.             | Create a specific order. 201 Created.      |
PATCH 	   | UPDATE              |                                  | Partial Update. 200 OK.                    |
DELETE 	   | DELETE              |                                  | Delete order. 200 OK.                      |


## Pluriel vs singulier

Nous utilisons le pluriel dans les noms de ressources :
* Collection de ressources : /v1/users
* Instance d’une ressource : /v1/users/007

## Casse des URI

Utilisation du spinal-case : 
Variante du snake case, le spinal case utilise des tirets courts « – » pour séparer les mots.
 
Avantages :
Permet aux compilateurs et/ou interprêteurs de le comprendre en tant que symbole unique, 
mais permettant au lecteur humain de séparer les mots de manière quasi naturelle

Inconvénients : certains langages de programmation ne peuvent l’accepter en tant que symbole 
(nom de variable ou fonction). 

Exemples : spinal-case, current-user, add-attribute-to-group, etc.

## Casse du body

Utilisation du snake_case

Exemple du contenu d'une réponse : 

```json
{
    "id": 91,
    "firstname": "one",
    "lastname": "Plus",
    "email": "one@gmail.com",
    "ssh_public_key": "ssh-rsa AAAAB__ONE",
    "pseudo": "oplus"
  }

```

## Versioning

Inscription de la version dans l'url : http://localhost:8090/v1/users/

La politique de maintenance des
 
# Non géré pour le moment

Les fonctionnalités ne sont pour l'instant pas gérées : 

* Pagination
* Filtres
* Tris
* Recherche



