package Tasks

import (
	"../TextIo"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"strconv"
)

type Patient struct {
	Surname string `json:"surname"`
	Name string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address string `json:"address"`
	CardId string `json:"card_id"`
	InsuranceId string `json:"insurance_id"`
}

type Patients []Patient;

func getPatient(surname string, name string, patronymic string, address string) Patient {
	cardId, _ := uuid.NewV4();
	insuranceId, _ := uuid.NewV4();

	return Patient{
		Surname: surname,
		Name: name,
		Patronymic: patronymic,
		Address: address,
		CardId: cardId.String(),
		InsuranceId: insuranceId.String(),
	};
}

func (patients Patients) toJson() string {
	res, _ := json.Marshal(patients);
	return string(res);
}

func fromJson(data string) Patients {
	bytes := []byte(data);
	var patients Patients;

	json.Unmarshal(bytes, &patients);
	return patients;
}

func findIndex(arr Patients, pred func(item Patient) bool) int {
	for i, item := range arr {
		if(pred(item)) {
			return i;
		}
	}
	return -1;
}

const file = "C:/Users/Олег/Desktop/patients.txt";

func Task2() {
	initPatients := make(Patients, 0);

 	initPatients = append(initPatients, getPatient("Kyrychenko", "Oleh", "Volodymyrovych", "Kharkiv"));
	initPatients = append(initPatients, getPatient("Nikolskiy", "Igor", "Bat'kovich", "Kharkiv"));
	initPatients = append(initPatients, getPatient("Syzrancev", "Maksym", "Bat'kovich", "Kharkiv"));

	str := initPatients.toJson();
	descr := TextIo.NewTextFile(file);
	descr.Write(str);

	content := descr.Read();
	fmt.Printf("ORIGINAL PATIENTS LIST: %s\n", content);

	patients := fromJson(content);
	cardIdToDelete := patients[2].CardId;

	indexToDelete := findIndex(patients, func(item Patient) bool {
		return Patient(item).CardId == cardIdToDelete;
	});
	patients = append(patients[:indexToDelete], patients[indexToDelete + 1:]...);
	descr.Write(patients.toJson());

	contentDeleted := descr.Read();
	fmt.Printf("PATIENTS LIST DELETED ITEM: %s\n", contentDeleted);
	deletedPatients := fromJson(contentDeleted);

	k := 3;
	prepended := make(Patients, k);
	for i := 0; i < k; i++ {
		prepended[i] = getPatient(
			"Test Surname" + strconv.Itoa(i),
			"Test Name " + strconv.Itoa(i),
			"Test Patronymic " + strconv.Itoa(i),
			"TEST address " + strconv.Itoa(i),
		);
	}

	prependedPatients := append(prepended, deletedPatients...);
	prependedJson := prependedPatients.toJson();
	descr.Write(prependedJson);

	contentPrepended := descr.Read();
	fmt.Printf("PATIENTS LIST PREPENDED ITEMS: %s\n", contentPrepended);
}