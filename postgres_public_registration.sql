create table registration
(
    id         bigserial not null
        constraint registration_pkey
            primary key,
    date_time  timestamp not null,
    id_patient bigint
        constraint registration_id_patient_fkey
            references patient,
    id_doctor  bigint
        constraint registration_id_doctor_fkey
            references doctor
);

alter table registration
    owner to postgres;

INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (45, '2021-02-10 18:00:00.000000', null, 1);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (46, '2021-02-10 11:00:00.000000', null, 2);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (50, '2021-02-10 15:00:00.000000', null, 2);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (51, '2021-02-10 16:00:00.000000', null, 2);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (52, '2021-02-10 17:00:00.000000', null, 2);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (53, '2021-02-10 18:00:00.000000', null, 2);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (54, '2021-02-10 11:00:00.000000', null, 3);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (55, '2021-02-10 12:00:00.000000', null, 3);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (58, '2021-02-10 15:00:00.000000', null, 3);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (59, '2021-02-10 16:00:00.000000', null, 3);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (61, '2021-02-10 18:00:00.000000', null, 3);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (62, '2021-02-10 11:00:00.000000', null, 4);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (63, '2021-02-10 12:00:00.000000', null, 4);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (66, '2021-02-10 15:00:00.000000', null, 4);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (67, '2021-02-10 16:00:00.000000', null, 4);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (68, '2021-02-10 17:00:00.000000', null, 4);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (69, '2021-02-10 18:00:00.000000', null, 4);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (70, '2021-02-10 11:00:00.000000', null, 5);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (71, '2021-02-10 12:00:00.000000', null, 5);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (74, '2021-02-10 15:00:00.000000', null, 5);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (76, '2021-02-10 17:00:00.000000', null, 5);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (60, '2021-02-10 17:00:00.000000', 51, 3);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (47, '2021-02-10 12:00:00.000000', 12, 2);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (43, '2021-02-10 16:00:00.000000', 5, 1);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (38, '2021-02-10 11:00:00.000000', 1, 1);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (39, '2021-02-10 12:00:00.000000', 1, 1);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (42, '2021-02-10 15:00:00.000000', 2, 1);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (75, '2021-02-10 16:00:00.000000', 2, 5);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (56, '2021-02-10 13:00:00.000000', 999, 3);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (64, '2021-02-10 13:00:00.000000', 999, 4);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (72, '2021-02-10 13:00:00.000000', 999, 5);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (48, '2021-02-10 13:00:00.000000', 999, 2);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (40, '2021-02-10 13:00:00.000000', 999, 1);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (78, '2021-02-10 13:00:00.000000', 999, null);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (73, '2021-02-10 14:00:00.000000', 999, 5);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (65, '2021-02-10 14:00:00.000000', 999, 4);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (49, '2021-02-10 14:00:00.000000', 999, 2);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (57, '2021-02-10 14:00:00.000000', 999, 3);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (41, '2021-02-10 14:00:00.000000', 567, 1);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (44, '2021-02-10 17:00:00.000000', 23, 1);
INSERT INTO public.registration (id, date_time, id_patient, id_doctor) VALUES (77, '2021-02-10 18:00:00.000000', 10, 5);