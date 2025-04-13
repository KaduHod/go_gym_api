package tests

import (
	"KaduHod/muscles_api/src/core"
	"KaduHod/muscles_api/src/database"
	repository "KaduHod/muscles_api/src/repositorys"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestRepositorys(t *testing.T) {
    if err := godotenv.Load("../.env"); err != nil {
        log.Fatal(err)
    }
	db := database.ConnetionMysql()
	defer db.Close()
	ammRepository := repository.AmmRepository{Db: db}

	tests := []struct {
		name    string
		filters map[string]string
		wantErr bool
	}{
		{
			name:    "no filters",
			filters: map[string]string{},
			wantErr: false,
		},
		{
			name: "with muscle group filter",
			filters: map[string]string{
				"muscle_group": "Chest",
			},
			wantErr: false,
		},
		{
			name: "with multiple filters",
			filters: map[string]string{
				"muscle_group": "Chest",
				"joint":        "Shoulder Glenohumeral",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ammRepository.GetAll(tt.filters)
			if (err != nil) != tt.wantErr {
				t.Errorf("AmmRepository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

    jointRepository := repository.JointRepository{Db: db}

    t.Run("Joints :: GetAll", func(t *testing.T) {
		joints, err := jointRepository.GetAll()
		if err != nil {
			t.Fatalf("GetAll failed: %v", err)
		}

		if len(joints) == 0 {
			t.Error("Expected at least one joint, got none")
		}

		for _, joint := range joints {
			if joint.Id == 0 {
				t.Error("Joint ID should not be zero")
			}
			if joint.Name == "" {
				t.Error("Joint name should not be empty")
			}
		}
	})

    t.Run("Joints :: GetById", func(t *testing.T) {
		// First get any joint to test with
		joints, err := jointRepository.GetAll()
		if err != nil || len(joints) == 0 {
			t.Fatal("Need at least one joint to test GetById")
		}

		testJoint := joints[0]
		joint, err := jointRepository.GetById(testJoint.Id)
		if err != nil {
			t.Fatalf("GetById failed: %v", err)
		}

		if joint.Id != testJoint.Id {
			t.Errorf("Expected joint ID %d, got %d", testJoint.Id, joint.Id)
		}

		if joint.Name != testJoint.Name {
			t.Errorf("Expected joint name %s, got %s", testJoint.Name, joint.Name)
		}

		// Test non-existent ID
		_, err = jointRepository.GetById(-1)
		if err == nil {
			t.Error("Expected error for non-existent ID, got nil")
		}
	})
	movementRepository := repository.MovementRepository{Db: db}
    t.Run("Moviment :: GetAll", func(t *testing.T) {
		movements, err := movementRepository.GetAll()
		if err != nil {
			t.Fatalf("GetAll failed: %v", err)
		}

		if len(movements) == 0 {
			t.Error("Expected at least one movement, got none")
		}

		for _, movement := range movements {
			if movement.Id == 0 {
				t.Error("Movement ID should not be zero")
			}
			if movement.Name == "" {
				t.Error("Movement name should not be empty")
			}
		}
	})

    t.Run("Moviment :: GetById", func(t *testing.T) {
		// First get any movement to test with
		movements, err := movementRepository.GetAll()
		if err != nil || len(movements) == 0 {
			t.Fatal("Need at least one movement to test GetById")
		}

		testMovement := movements[0]
		movement, err := movementRepository.GetById(testMovement.Id)
		if err != nil {
			t.Fatalf("GetById failed: %v", err)
		}

		if movement.Id != testMovement.Id {
			t.Errorf("Expected movement ID %d, got %d", testMovement.Id, movement.Id)
		}

		if movement.Name != testMovement.Name {
			t.Errorf("Expected movement name %s, got %s", testMovement.Name, movement.Name)
		}

		// Test non-existent ID
		_, err = movementRepository.GetById(-1)
		if err == nil {
			t.Error("Expected error for non-existent ID, got nil")
		}
	})

	muscleRepository := repository.MuscleRepository{Db: db}

	t.Run("Muscle :: GetAll", func(t *testing.T) {
		groups, err := muscleRepository.GetAll()
		if err != nil {
			t.Fatalf("GetAll failed: %v", err)
		}

		if len(groups) == 0 {
			t.Error("Expected at least one muscle group, got none")
		}

		for _, group := range groups {
			if group.Id == nil {
				t.Error("Muscle group ID should not be nil")
			}
			if group.Name == "" {
				t.Error("Muscle group name should not be empty")
			}
		}
	})

	t.Run("Muscle :: GetById", func(t *testing.T) {
		groups, err := muscleRepository.GetAll()
		if err != nil || len(groups) == 0 {
			t.Fatal("Need at least one muscle group to test GetById")
		}

		testGroup := groups[0]
		group, err := muscleRepository.GetById(*testGroup.Id)
		if err != nil {
			t.Fatalf("GetById failed: %v", err)
		}

		if *group.Id != *testGroup.Id {
			t.Errorf("Expected group ID %d, got %d", *testGroup.Id, *group.Id)
		}

		if group.Name != testGroup.Name {
			t.Errorf("Expected group name %s, got %s", testGroup.Name, group.Name)
		}
	})

	t.Run("Muscle :: GetWithPortions", func(t *testing.T) {
		groups, err := muscleRepository.GetWithPortions()
		if err != nil {
			t.Fatalf("GetWithPortions failed: %v", err)
		}

		if len(*groups) == 0 {
			t.Error("Expected at least one muscle group with portions")
		}

		for _, group := range *groups {
			if len(group.Portions) == 0 {
				t.Errorf("Expected portions for group %s, got none", group.Name)
			}
		}
	})

	t.Run("Muscle :: GetAllPortions", func(t *testing.T) {
		portions, err := muscleRepository.GetAllPortions()
		if err != nil {
			t.Fatalf("GetAllPortions failed: %v", err)
		}

		if len(portions) == 0 {
			t.Error("Expected at least one muscle portion")
		}

		for _, portion := range portions {
			if portion.Id == nil {
				t.Error("Portion ID should not be nil")
			}
			if portion.Name == "" {
				t.Error("Portion name should not be empty")
			}
			if portion.MuscleGroupId == nil {
				t.Error("MuscleGroupId should not be nil")
			}
		}
	})

	t.Run("Muscle :: GetPortionById", func(t *testing.T) {
		portions, err := muscleRepository.GetAllPortions()
		if err != nil || len(portions) == 0 {
			t.Fatal("Need at least one portion to test GetPortionById")
		}

		testPortion := portions[0]
		portion, err := muscleRepository.GetPortionById(*testPortion.Id)
		if err != nil {
			t.Fatalf("GetPortionById failed: %v", err)
		}

		if *portion.Id != *testPortion.Id {
			t.Errorf("Expected portion ID %d, got %d", *testPortion.Id, *portion.Id)
		}
	})
    tokenRepository := repository.TokenRepository{Db: db}
    t.Run("Token :: List tokens", func(t *testing.T) {
        _, err := tokenRepository.GetTokens(core.ApiUser{Id: 1})
        if err != nil {
            t.Fatal(err)
        }
    })
    t.Run("Token :: Get Tokens by login", func(t *testing.T) {
        _, err := tokenRepository.GetTokensByLogin("KaduHod")
        if err != nil {
            t.Fatal(err)
        }
    })
    exerciseRepository := repository.ExerciseRepository{Db: db}
    t.Run("Exercise :: Get all exercises", func(t *testing.T) {
        _, err := exerciseRepository.GetExercises()
        if err != nil {
            t.Fatal(err)
        }
    })
    t.Run("Exercise :: Get details", func(t *testing.T) {
        _, err := exerciseRepository.GetExerciseDetails(84)
        if err != nil {
            t.Fatal(err)
        }
    })
}
