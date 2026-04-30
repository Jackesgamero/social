package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/jackesgamero/social/internal/store"
)

var usernames = []string{
	"Alejandro", "Beatriz", "Carlos", "Daniela", "Eduardo",
	"Fernanda", "Gabriel", "Helena", "Ignacio", "Javier",
	"Karla", "Luis", "Mariana", "Nicolás", "Olivia",
	"Pablo", "Raquel", "Santiago", "Teresa", "Uriel",
	"Valeria", "Walter", "Ximena", "Yolanda", "Zoe",
	"Adrián", "Blanca", "Cristian", "Diana", "Esteban",
	"Fabiola", "Guillermo", "Hugo", "Isabel", "Jorge",
	"Kevin", "Lorena", "Miguel", "Natalia", "Oscar",
	"Patricia", "Ricardo", "Sofía", "Tomás", "Úrsula",
	"Vicente", "Wendy", "Xavier", "Yisus", "Zacarías",
}

var titles = []string{
	"Lo que nadie te cuenta",
	"Productividad al máximo",
	"El error que todos cometen",
	"Mi rutina de mañana",
	"¿Team café o team té?",
	"3 Tips para tu carrera",
	"Mentalidad de éxito",
	"Lo aprendí a las malas",
	"Herramientas que uso hoy",
	"POV: Es lunes otra vez",
	"Guía rápida para principiantes",
	"La verdad sobre el éxito",
	"Mis libros favoritos",
	"¿Cómo lo logré?",
	"Menos es más",
	"Unboxing inesperado",
	"Hack de organización",
	"Preguntas y respuestas",
	"Detrás de cámaras",
	"Hablemos de esto...",
}

var contents = []string{
	"Si solo pudieras cambiar una cosa hoy, que fuera esta.",
	"Llevo meses probando esto y los resultados me volaron la cabeza.",
	"Parece obvio, pero el 90 porciento de la gente lo ignora por completo.",
	"Hoy quiero ser totalmente honesto con ustedes sobre lo que pasó.",
	"Si estás esperando una señal para empezar, es esta.",
	"Guardá este post, porque lo vas a necesitar más tarde.",
	"Todo cambió cuando dejé de enfocarme en los resultados.",
	"Hablemos de esa verdad incómoda que nadie quiere admitir.",
	"No necesitas más tiempo, necesitas mejores prioridades.",
	"Esta es la pregunta que más me hicieron durante la semana.",
	"La mayoría de los consejos que escuchas sobre esto están mal.",
	"Hace un año estaba en un lugar totalmente distinto al de hoy.",
	"Te comparto el sistema exacto que uso para organizar mi día.",
	"¿Alguna vez sentiste que estás trabajando mucho pero no avanzas?",
	"Acá te dejo 3 herramientas que me ahorran horas de trabajo.",
	"No se trata de ser el mejor, se trata de ser constante.",
	"Si estás leyendo esto, probablemente necesites un descanso.",
	"Lo que ves en redes es solo el 5 porciento del esfuerzo real.",
	"Esto es lo que desearía haber sabido cuando empecé.",
	"Basta de poner excusas, el momento perfecto no existe.",
}

var tags = []string{
	"#productividad", "#emprendimiento", "#marketingdigital", "#desarrollopersonal",
	"#motivacion", "#creadordecontenido", "#exito", "#redessociales",
	"#mindset", "#trabajoremoto", "#tips", "#estilodevida",
	"#metas", "#creatividad", "#aprendizaje", "#negocios",
	"#comunidad", "#inspiracion", "#hack", "#growth",
}

var comments = []string{
	"¡Totalmente de acuerdo! Justo lo que necesitaba leer hoy.",
	"¿Podrías dar más detalles sobre la herramienta que mencionas?",
	"Guardado. Esto me va a servir muchísimo para la semana que viene.",
	"Me pasó exactamente lo mismo hace un par de meses, ¡es tal cual!",
	"Excelente contenido, gracias por compartir con tanta claridad.",
	"No lo había pensado de esa manera, me cambiaste la perspectiva.",
	"¡Qué buen hook! Me quedé leyendo hasta el final.",
	"100% real. El 95 porciento del esfuerzo no se ve en las fotos.",
	"¿Tenés algún libro o curso que recomiendes sobre este tema?",
	"¡Gracias por la honestidad! Hace falta más contenido así.",
	"Directo al grano, como me gusta. ¡Seguí así!",
	"Justo estaba por tirar la toalla y apareció este post. Gracias.",
	"Ojalá hubiera sabido esto cuando empecé mi proyecto.",
	"Muy buen tip, lo voy a poner en práctica mañana mismo.",
	"¿Para cuándo un vivo hablando más sobre esto?",
	"Etiqueto a @amigo para que vea esto, le va a encantar.",
	"Amé el diseño del post y la info está de 10.",
	"La constancia es la clave, no hay otra vuelta que darle.",
	"¡Qué buen aporte! Sumaría también el tema de la disciplina.",
	"Totalmente, el momento perfecto es hoy. ¡A darle con todo!",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
