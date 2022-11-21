
create table coach (
    "id" serial primary key,
    "username" varchar(30) not null,
    "password" varchar(100) not null
);

create table "workout" (
    "id" serial primary key,
	"coach_id" int not null,
    "name" varchar(50) not null,
    "description" varchar(100) not null,
    constraint fk_coach foreign key(coach_id) references coach(id)
); 

create table "video" (
    "id" serial primary key,
    "link" varchar(500) not null
);

create table "video_to_workout" (
    "workout_id" int not null,
    "video_id" int not null,
    constraint fk_workout foreign key(workout_id) references workout(id),
    constraint fk_video foreign key(video_id) references video(id)
);

-- test / test
insert into coach("id", "username", "password") values 
(1, 'test', '$2b$10$CCAFoZOhJA7KtzZoX2B9u.a3fWlIIg/fSMm3I5VWHdW5A26zO7RHG');

insert into workout ("coach_id", "name", "description") values 
(1, 'Workout 1', 'Workout 1 Description');

insert into video ("id", "link") values 
(1, 'https://stackoverflow.com/questions/18389124/simulate-create-database-if-not-exists-for-postgresql'),
(2, 'https://www.calhoun.io/6-tips-for-using-strings-in-go/#:~:text=Multiline%20strings,%60');

insert into video_to_workout ("workout_id", "video_id") values
(1, 1),
(1, 2);