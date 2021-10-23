package services

import (
	"errors"
	"sort"

	"go-kbs-notification/entities/models"
	"go-kbs-notification/repositories"
	"go-kbs-notification/services/gateway"
)

// for testing
var notifCount = 0

type smsNotificationService struct {
	repository repositories.FeatureFlagRepository
}

func NewSMSNotificationService(repository repositories.FeatureFlagRepository) SMSNotificationService {
	return &smsNotificationService{
		repository: repository,
	}
}

func (s *smsNotificationService) SendSMS(req models.SMS) error {
	notifCount += 1
	s.ForTestingFlag()

	gateway, err := s.ProcessSMSNotificationGateway()
	if err != nil {
		return err
	}

	err = gateway.SendSMS(req.PhoneNumber, req.Message)
	if err != nil {
		return err
	}

	return nil
}

func (s *smsNotificationService) ProcessSMSNotificationGateway() (gateway.SMSGateway, error) {
	featureFlag, err := s.repository.Get()
	if err != nil {
		return nil, err
	}
	s.SortFlag(featureFlag)

	for _, flag := range featureFlag {
		if flag.Active {
			switch flag.Vendor {
			case "twilio":
				return gateway.NewTwilioGateway(), nil
			case "nexmo":
				return gateway.NewNexmoGateway(), nil
			case "vendor3":
				return gateway.NewVendor3Gateway(), nil
			default:
				return nil, errors.New("vendor not found")
			}
		}
	}

	return nil, errors.New("flag not found")
}

func (s *smsNotificationService) SortFlag(flags []models.FeatureFlag) {

	sort.Slice(flags, func(i, j int) bool {
		var sortedByActice, sortedByLowerPrice bool

		// sort by active
		sortedByActice = flags[i].Active && !flags[j].Active

		// sort by lowest sold price
		if flags[i].Active == flags[j].Active {
			sortedByLowerPrice = flags[i].Price < flags[j].Price
			return sortedByLowerPrice
		}
		return sortedByActice
	})
}

func (s *smsNotificationService) ForTestingFlag() {
	if notifCount%2 == 0 && notifCount%3 == 0 {
		s.repository.Store([]models.FeatureFlag{
			{
				Vendor: "twilio",
				Price:  100,
				Active: false,
			},
			{
				Vendor: "nexmo",
				Price:  200,
				Active: false,
			},
			{
				Vendor: "vendor3",
				Price:  300,
				Active: true,
			},
		})
		return
	} else if notifCount%2 == 0 {
		s.repository.Store([]models.FeatureFlag{
			{
				Vendor: "twilio",
				Price:  100,
				Active: false,
			},
			{
				Vendor: "nexmo",
				Price:  200,
				Active: true,
			},
			{
				Vendor: "vendor3",
				Price:  300,
				Active: true,
			},
		})
		return
	} else if notifCount%3 == 0 {
		s.repository.Store([]models.FeatureFlag{
			{
				Vendor: "twilio",
				Price:  100,
				Active: true,
			},
			{
				Vendor: "nexmo",
				Price:  200,
				Active: false,
			},
			{
				Vendor: "vendor3",
				Price:  300,
				Active: true,
			},
		})
		return
	}

	s.repository.Store([]models.FeatureFlag{
		{
			Vendor: "twilio",
			Price:  100,
			Active: true,
		},
		{
			Vendor: "nexmo",
			Price:  200,
			Active: true,
		},
		{
			Vendor: "vendor3",
			Price:  300,
			Active: true,
		},
	})
}
