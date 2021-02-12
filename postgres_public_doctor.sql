create table doctor
(
    id         bigserial not null
        constraint doctor_pkey
            primary key,
    first_name varchar(50),
    last_name  varchar(50),
    email      varchar(50),
    gender     varchar(50),
    password   varchar(50)
);

alter table doctor
    owner to postgres;

INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (1, 'Konstanze', 'Brachell', 'kbrachell0@tuttocitta.it', 'Bigender', 'bUvfIGKy');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (2, 'Pia', 'Mardy', 'pmardy1@ezinearticles.com', 'Non-binary', 'aKupSln');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (3, 'Micky', 'Meiner', 'mmeiner2@tinypic.com', 'Genderfluid', 'Y6KcJ2SjFTdL');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (4, 'Davidson', 'Yuille', 'dyuille3@digg.com', 'Male', '9LbyvJXZH');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (5, 'Dani', 'Sporner', 'dsporner4@devhub.com', 'Non-binary', 'nMq6T5Su');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (6, 'Cecilio', 'Muffin', 'cmuffin5@lycos.com', 'Female', 'oM5UIdaDM');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (7, 'Kathie', 'Shrubsall', 'kshrubsall6@gnu.org', 'Female', 'b9VBWQ');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (8, 'Talia', 'Hare', 'thare7@tumblr.com', 'Agender', 'BtqPa3U');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (9, 'Osbert', 'Pattillo', 'opattillo8@google.ca', 'Genderfluid', 'KYSzPmrYhwh');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (10, 'Ferdy', 'Wagen', 'fwagen9@narod.ru', 'Genderfluid', 'N83CgOzC');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (11, 'Manolo', 'Blazejewski', 'mblazejewskia@usda.gov', 'Polygender', 's1ihx70kdz8');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (12, 'Marieann', 'Tresvina', 'mtresvinab@cpanel.net', 'Bigender', 'dIJB9uUNcFj');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (13, 'Wilburt', 'Zwicker', 'wzwickerc@china.com.cn', 'Bigender', 'EXROG73');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (14, 'Viviene', 'Cunliffe', 'vcunliffed@cnbc.com', 'Polygender', 't09ydiXSu');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (15, 'Ben', 'Hayler', 'bhaylere@nydailynews.com', 'Genderqueer', 'CU2l77S');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (16, 'Rani', 'Coyett', 'rcoyettf@nifty.com', 'Genderqueer', 'Y5QgaQKy');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (17, 'Klement', 'Buzine', 'kbuzineg@kickstarter.com', 'Male', 'CXBpqZr5h');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (18, 'Malchy', 'Lugton', 'mlugtonh@toplist.cz', 'Female', '4dN7TFLL');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (19, 'Timothee', 'Easby', 'teasbyi@elpais.com', 'Non-binary', 'xSuvdKg');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (20, 'Husein', 'Lindenman', 'hlindenmanj@parallels.com', 'Genderfluid', 'mg5IXytnrm');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (21, 'Reinaldos', 'Allsobrook', 'rallsobrookk@zdnet.com', 'Polygender', '0IKWssaRhkMC');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (22, 'Evvy', 'Dorkins', 'edorkinsl@typepad.com', 'Female', 'pQvns6');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (23, 'Norean', 'McManamon', 'nmcmanamonm@blog.com', 'Bigender', 'VaUsquNPK9');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (24, 'Lyell', 'Gillan', 'lgillann@un.org', 'Male', 'H72UQqglkTB');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (25, 'Lock', 'McCraine', 'lmccraineo@upenn.edu', 'Female', 'ak0G7eeL');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (26, 'Syd', 'Worrell', 'sworrellp@usatoday.com', 'Male', '0STzBRxvDClJ');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (27, 'Pat', 'Mangeot', 'pmangeotq@google.fr', 'Female', 'wl4YqZpnCll');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (28, 'Shalne', 'Greatrakes', 'sgreatrakesr@uiuc.edu', 'Genderfluid', '42BtmfCruID');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (29, 'Kristos', 'Jelkes', 'kjelkess@booking.com', 'Polygender', 'OmgIComDlNZ');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (30, 'Christopher', 'Negro', 'cnegrot@dion.ne.jp', 'Non-binary', 'GCK2T4dan');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (31, 'Phip', 'Merrywether', 'pmerrywetheru@twitpic.com', 'Agender', 'ePONPCZ0SXPG');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (32, 'Briano', 'Mecco', 'bmeccov@booking.com', 'Genderfluid', 'Y2ea5kb3uqd1');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (33, 'Nicol', 'Risby', 'nrisbyw@hc360.com', 'Bigender', 'gEPNMHUZk4');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (34, 'Penny', 'Doggart', 'pdoggartx@cnbc.com', 'Female', '3rHhFja');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (35, 'Charisse', 'Vevers', 'cveversy@jigsy.com', 'Polygender', 't457KmyhSQA');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (36, 'Ema', 'Pattrick', 'epattrickz@com.com', 'Agender', 'E05kkS');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (37, 'Shawna', 'Signoret', 'ssignoret10@state.tx.us', 'Genderqueer', 'CjqNmhH');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (38, 'Hughie', 'Keerl', 'hkeerl11@skyrock.com', 'Bigender', 'bqme19');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (39, 'Drona', 'Molson', 'dmolson12@oaic.gov.au', 'Agender', 'QvyJKOBn0rI');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (40, 'Rhys', 'Manna', 'rmanna13@fema.gov', 'Polygender', 'NnFOX3j9Eea');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (41, 'Chelsy', 'Povele', 'cpovele14@disqus.com', 'Genderqueer', 'Z9RuIeeNy15');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (42, 'Elga', 'Fratson', 'efratson15@telegraph.co.uk', 'Male', '8BvI5U');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (43, 'Tonnie', 'Towers', 'ttowers16@skype.com', 'Male', 'An2hYBRWf3FD');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (44, 'Salem', 'Iashvili', 'siashvili17@intel.com', 'Non-binary', 'UOiPUoNt');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (45, 'Hernando', 'Spiniello', 'hspiniello18@arstechnica.com', 'Genderfluid', '02FPhoFH');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (46, 'Hamlen', 'Koenen', 'hkoenen19@google.nl', 'Non-binary', 'mywaSm');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (47, 'Brand', 'Lackney', 'blackney1a@narod.ru', 'Polygender', 'pOAZTTqiLSt');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (48, 'Halimeda', 'Holbie', 'hholbie1b@netvibes.com', 'Polygender', 'rmO0MhFG');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (49, 'Diego', 'Mose', 'dmose1c@ibm.com', 'Male', 'oUTuwam');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (50, 'Devin', 'Pittman', 'dpittman1d@slideshare.net', 'Bigender', '4QLcBU');
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (53, 'Max', 'Vasil', null, null, null);
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (54, 'Max', 'Vasil', null, null, null);
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (55, 'Max', 'Vasil', null, null, null);
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (56, 'Max', 'Vasil', null, null, null);
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (57, 'Max', 'Vasil', null, null, null);
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (51, 'jopa', 'Vasil', null, null, null);
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (59, 'Mishaaaaaaaaaaaaaaaaaaaa', 'Eblaneeee', null, null, null);
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (60, 'Mishaaaaaaaaaaaaaaaaaaaa', 'Eblaneeee', null, null, null);
INSERT INTO public.doctor (id, first_name, last_name, email, gender, password) VALUES (61, 'Jasleen', 'Makkar', null, null, null);