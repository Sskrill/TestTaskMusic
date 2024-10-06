Зайдите в makefile и напишите свои параметры для входа в бд postgres 

Введите команду **make run**

    	"song/add" - Добавление песни (добавление в json формате в body запроса)
     
		"song/edit/{id}" - Изменение песни по id 
  
		"song/details/{song_name}/{performer_name}" - Получение деталей песни по ее названию и название исполнителя
  
		"song/delete/{id}" - Удаление песни по id 
  
		"song/text/{song_name}/{performer_name}" Получение текста песни по ее названию и название исполнителя
  
		"song/filters" - Поиск песен по фильтру (фильтры в json формате в body запроса)
